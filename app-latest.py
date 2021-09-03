
'''
#############################  wunderDB  #############################

wunderDB is a JSON-based micro Document DB, inspired by MongoDB.

version 0.1 Beta
developed by Tanmoy Sen Gupta

* Get started by downloading the app.py and db.json files.
* Create a Python Virtual Environment.
* Install Flask and shortuuid using pip.
* Run 'flask run' and your DB is ready to use!

When the development server starts access the instructions by going to localhost:5000/get-started

'''

import hashlib, json, string, random, secrets, shortuuid
from flask import Flask, request, jsonify, redirect, send_from_directory
from os import path, write
from flask_cors import CORS
from datetime import datetime
import smtplib, ssl, time, hashlib
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart

app = Flask(__name__)

CORS(app)

'''@app.route('/')
def index():
    #return app.send_static_file('../index.html')
    #return send_from_directory('./index.html')
    return "Hellow"
    
@app.route('/wdash')
def wdash():
    # return redirect('https://wdb.tanmoysg.com/wdash/index.html')
    return send_from_directory('../wdash/index.html')
    # return app.send_static_file('../wdash/index.html')'''


server_credentials = json.load(open('server-config.json'))

port = server_credentials['port']
serverAddress = server_credentials['mail-server']
sender = server_credentials['sender']
password = server_credentials['password']

def write_json(data, filename):
    with open(filename, 'w') as f:
        json.dump(data, f, indent=4)
        
def hash_otp(mail, otp, timestamp):
    complete_string = mail+"."+otp+"."+timestamp
    return hashlib.md5(complete_string.encode()).hexdigest()
        
def generate_OTP():
    return ''.join(random.choices(string.ascii_uppercase + string.ascii_lowercase + string.digits , k = 6))

def send_mail(recipient, mail_template, content):
    message = MIMEMultipart("alternative")
    message["Subject"] = content["subject"]
    message["From"] = sender
    message["To"] = recipient
    template = open(mail_template).read().replace(content["replacement_token"], content["replacement_string"])
    html = MIMEText(template, "html")
    message.attach(html)
    context = ssl.create_default_context()
    with smtplib.SMTP_SSL(serverAddress, port, context=context) as server:
        server.login(sender, password)
        server.sendmail(sender, recipient, message.as_string())


def generate_mail(recipient, mode, payload="none"):
    if mode=="generate-otp-mail":
        content = {
            "subject" : "Verify you wunderDB account",
            "replacement_token": "%otp%",
            "replacement_string": payload["otp"]
        }

        send_mail(recipient, "otp_generation_template.txt", content)

    elif mode=="verified-otp-mail":
        content = {
            "subject" : "Account Verified.",
            "replacement_token": "%user%",
            "replacement_string": recipient
        }
        send_mail(recipient, "otp_verified_template.txt", content)


def create_new_request(mail, otp):
    with open("authlib.json") as authlib:
        authdict = json.load(authlib)
        if mail not in authdict.keys():
            ts = time.time()
            timestamp = datetime.fromtimestamp(ts).strftime('%Y-%m-%d %H:%M:%S')
            authdict[mail] = {
                "mail" : mail,
                "timestampLatestOTP": timestamp ,
                "latestOTP": hash_otp(mail, otp, timestamp),
                "status" : "unverified"
            }
            write_json(authdict, "authlib.json")
            generate_mail(mail, "generate-otp-mail", {"otp": otp})
            return "Created"
        else:
            generate_mail(mail, "generate-otp-mail", {"otp": authdict[mail]["latestOTP"]})
            print("Request Exists")

@app.route("/register/authenticate", methods = ['POST'])
def initialize_new_authentication():
    incoming_data = request.get_json(force = True)
    with open("users.json", 'r') as users:
        user = json.load(users)
        users_list = user['users']
        if incoming_data['username'] not in users_list.keys():    
            otp = generate_OTP()
            create_new_request(incoming_data['username'], otp)
            return jsonify({ 
                "status_code" : '1',
                "response" : "OTP Sent." 
            })
        else:
            return jsonify({ 
                "status_code" : '1',
                "response" : "User Exists." 
            })

@app.route("/register/verify", methods = ['POST'])
def authenticate_OTP():
    incoming_data = request.get_json(force = True)
    mail = incoming_data['username']
    otp = incoming_data['otp']
    with open("authlib.json") as authlib:
        temp = json.load(authlib)
        if mail in temp.keys():
            if temp[mail]["latestOTP"] == hash_otp(mail, otp, temp[mail]["timestampLatestOTP"]):
                temp.pop(mail)
                write_json(temp, "authlib.json")
                generate_mail(mail, "verified-otp-mail")
                print("Verified")
                return register(incoming_data)
            else:
                print("Wrong OTP")
                return jsonify({ 
                    "status_code" : '0',
                    "response" : "Wrong OTP." 
                })
        else:
            return jsonify({ 
                "status_code" : '0',
                "response" : "User already exists." 
            })


# GETTING STARTED
@app.route('/get-started', methods = ['GET'])
def get_started():
    with open('templates.json') as templates:
        instructions = json.load(templates)
        return jsonify({
            "_00_message" : "Welcome to wunderDB !",
            "_01_about"   : "wunderDB is a JSON-based micro Document DB inspired by MongoDB.",
            "_02_version" : "0.1 Beta",
            "_03_creator" : "Tanmoy Sen Gupta",
            "_04_instructions" : instructions
        })


# CREATE CLUSTER / REGISTER
# @app.route('/register', methods = ['POST'])
def register(user_data):
    with open("users.json", 'r') as users:
        user = json.load(users)
        users_list = user['users']
        if user_data['username'] not in users_list.keys():
            salt = user_data['username'].encode('utf-8')
            hashed_password = hashlib.sha512(user_data['password'].encode('utf-8') + salt).hexdigest()

            tokens = []
            cluster_id =shortuuid.uuid()
            for i in range(3):
                tokens.append(secrets.token_hex(16))
            
            user_init = {
                "_cluster_id" : cluster_id,
                "name" : user_data['name'],
                "username" : user_data['username'],
                "password" : hashed_password,
                "access_tokens": tokens
            }

            users_list[user_data['username']] = user_init
            write_json(user, 'users.json')

            cluster_data = {
                "_cluster_id": cluster_id,
                "access_tokens": tokens,
                "databases":{}
            }

            cluster_init = {
                "clusters" : {}
            }

            file = "./clusters/"+cluster_id+".json"

            with open(file , 'w+' ) as json_file:
                write_json(cluster_init, file)
                data = json.load(json_file)
                temp = data['clusters']
                temp[cluster_id] = cluster_data
                write_json(data, file)
            return jsonify({ 
                "status_code" : '1',
                "response" : "Cluster Created",
                "cluster_id" : cluster_id , 
                "access_tokens" : tokens
                })
        else:
            return jsonify({ 
                "status_code" : '0',
                "response" : "User already exists." 
                
            })

# LOGIN
@app.route('/login', methods = ['POST'])
def login():
    user_data = request.get_json(force = True)
    file = "users.json"
    with open(file) as user_list:
        users = json.load(user_list)
        user = users['users']
        if user_data["username"] in user.keys():
            salt = user_data['username'].encode('utf-8')
            hashed_password = hashlib.sha512(user_data['password'].encode('utf-8') + salt).hexdigest()
            if user[user_data["username"]]['username'] == user_data['username'] and user[user_data["username"]]['password'] == hashed_password :
                return jsonify({
                    "status_code" : '1',
                    "response" : "Logged-in",
                    "name" : user[user_data["username"]]['name'],
                    "cluster_id" : user[user_data["username"]]['_cluster_id'],
                    "access_token" : random.choice(user[user_data["username"]]['access_tokens'])
                })
            else:
                return jsonify({
                    "status_code" : '0',
                    "response" : "Password is wrong."
            })
        else:
            return jsonify({
                "status_code" : '0',
                "response" : "User doesn't exist"
            })

# CREATE DATABASE
def create_database(cluster_id , payload):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        db_json = payload

        database_id =shortuuid.uuid()
        db_data={
            "_uuid": database_id,
            "db_name": db_json['name'],
            "collections": {}
        }
        temp = cluster[cluster_id]['databases']
        temp[db_json['name']] = db_data
        write_json(clusters, file)
        return jsonify({
                "status_code" : '1',
                "response" : "Database Created with ID "+database_id
            })


# CREATE COLLECTION
def create_collection(cluster_id , payload):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = payload
        if cluster_json['database'] in cluster[cluster_id]["databases"].keys():
            collection_id =shortuuid.uuid()
            schema = cluster_json['schema']
            schema.update({"_id" : "ID"})
            collection_data={
                "_uuid": collection_id,
                "collection_name": cluster_json['name'],
                "schema": schema,
                "data": {}
                }
            temp = cluster[cluster_id]['databases'][cluster_json['database']]['collections']
            temp[cluster_json['name']] = collection_data
            write_json(clusters, file)
            return jsonify({
                        "status_code" : '1',
                        "response" : "Collection created with ID "+ collection_id
                    })
        else:
            return jsonify({
                    "status_code" : '0',
                    "response" : "Database doesn't exist."
                })


# DELETE COLLECTION
def delete_collection(cluster_id , payload):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = payload
        if cluster_json['database'] in cluster[cluster_id]["databases"].keys():
            database = cluster[cluster_id]["databases"][cluster_json['database']]
            if cluster_json['collection'] in cluster[cluster_id]["databases"][cluster_json['database']]["collections"].keys():
                database["collections"].pop(cluster_json["collection"])
                write_json(clusters, file)
                return jsonify({
                            "status_code" : '1',
                            "response" : "Collection Deleted"
                        })
            else:
                return jsonify({
                    "status_code" : '0',
                    "response" : "Collection doesn't exist."
                })
        else:
            return jsonify({
                    "status_code" : '0',
                    "response" : "Database doesn't exist."
                })
                

# CLONE ALL
def clone_all_collections(cluster_id , payload):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = payload
        if cluster_json['sourceDatabase'] in cluster[cluster_id]["databases"].keys():
            if cluster_json['destinationDatabase'] in cluster[cluster_id]["databases"].keys():
                destinationDatabase = cluster[cluster_id]["databases"][cluster_json['destinationDatabase']]
                sourceDatabase = cluster[cluster_id]["databases"][cluster_json['sourceDatabase']]
                for collectionName in cluster[cluster_id]["databases"][cluster_json['sourceDatabase']]["collections"].keys():
                    tempColName = "cloned-"+collectionName+"@"+cluster_json['sourceDatabase']
                    destinationDatabase["collections"][tempColName] = {}
                    destinationDatabase["collections"][tempColName]['_uuid'] = shortuuid.uuid()
                    destinationDatabase["collections"][tempColName]['collection_name'] = tempColName
                    destinationDatabase["collections"][tempColName]['schema'] =  sourceDatabase["collections"][collectionName].get('schema')
                    destinationDatabase["collections"][tempColName]['data'] =  sourceDatabase["collections"][collectionName].get('data')
                    write_json(clusters, file)
                write_json(clusters, file)
                return jsonify({
                            "status_code" : '1',
                            "response" : "Collections Cloned."
                        })
            else:
                return jsonify({
                    "status_code" : '0',
                    "response" : "Destination Database doesn't exist."
                })
        else:
            return jsonify({
                    "status_code" : '0',
                    "response" : "Source Database doesn't exist."
                })
                


# CLONE COLLECTION
def clone_collection(cluster_id , payload):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = payload
        if cluster_json['collection'] == "clone-all":
            return clone_all_collections(cluster_id, payload)
        if cluster_json['sourceDatabase'] in cluster[cluster_id]["databases"].keys():
            if cluster_json['destinationDatabase'] in cluster[cluster_id]["databases"].keys():
                sourceDatabase = cluster[cluster_id]["databases"][cluster_json['sourceDatabase']]
                destinationDatabase = cluster[cluster_id]["databases"][cluster_json['destinationDatabase']]
                if cluster_json['collection'] in cluster[cluster_id]["databases"][cluster_json['sourceDatabase']]["collections"].keys():
                    if cluster_json['collection'] in cluster[cluster_id]["databases"][cluster_json["destinationDatabase"]]["collections"].keys():
                        if cluster_json['actionIfExists'] == "replace":
                            destinationDatabase["collections"].pop(cluster_json["collection"])
                            destinationDatabase["collections"][cluster_json['collection']] = {}
                            destinationDatabase["collections"][cluster_json['collection']]['_uuid'] = shortuuid.uuid()
                            destinationDatabase["collections"][cluster_json['collection']]['collection_name'] = sourceDatabase["collections"][cluster_json['collection']].get('collection_name')
                            destinationDatabase["collections"][cluster_json['collection']]['schema'] = sourceDatabase["collections"][cluster_json['collection']].get('schema')
                            destinationDatabase["collections"][cluster_json['collection']]['data'] = sourceDatabase["collections"][cluster_json['collection']].get('data')
                            write_json(clusters, file)
                            return jsonify({
                                        "status_code" : '1',
                                        "response" : "Collection Copied!"
                                    })
                        elif cluster_json['actionIfExists'] == "rename":
                            destinationDatabase["collections"][cluster_json['newName']] = {}
                            destinationDatabase["collections"][cluster_json['newName']]['_uuid'] = shortuuid.uuid()
                            destinationDatabase["collections"][cluster_json['newName']]['collection_name'] = cluster_json['newName']
                            destinationDatabase["collections"][cluster_json['newName']]['schema'] = sourceDatabase["collections"][cluster_json['collection']].get('schema')
                            destinationDatabase["collections"][cluster_json['newName']]['data'] = sourceDatabase["collections"][cluster_json['collection']].get('data')
                            write_json(clusters, file)
                            return jsonify({
                                        "status_code" : '1',
                                        "response" : "Collection Copied!"
                                    })
                        elif  cluster_json['actionIfExists'] == "abort":
                            return jsonify({
                                        "status_code" : '0',
                                        "response" : "Process aborted as collection with same name exists in destination database."
                                    })
                    else:
                        destinationDatabase["collections"][cluster_json['collection']] = {}
                        destinationDatabase["collections"][cluster_json['collection']]['_uuid'] = shortuuid.uuid()
                        destinationDatabase["collections"][cluster_json['collection']]['collection_name'] = sourceDatabase["collections"][cluster_json['collection']].get('collection_name')
                        destinationDatabase["collections"][cluster_json['collection']]['schema'] = sourceDatabase["collections"][cluster_json['collection']].get('schema')
                        destinationDatabase["collections"][cluster_json['collection']]['data'] = sourceDatabase["collections"][cluster_json['collection']].get('data')
                        write_json(clusters, file)
                        return jsonify({
                                    "status_code" : '1',
                                    "response" : "Collection Copied!"
                                })
                else:
                    return jsonify({
                        "status_code" : '0',
                        "response" : "Collection doesn't exist."
                    })
            else:
                return jsonify({
                    "status_code" : '0',
                    "response" : "Destination Database doesn't exist."
                }) 
        else:
            return jsonify({
                    "status_code" : '0',
                    "response" : "Source Database doesn't exist."
                })


# MIGRATE ALL
def migrate_all_collections(cluster_id , payload):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = payload
        if cluster_json['sourceDatabase'] in cluster[cluster_id]["databases"].keys():
            if cluster_json['destinationDatabase'] in cluster[cluster_id]["databases"].keys():
                destinationDatabase = cluster[cluster_id]["databases"][cluster_json['destinationDatabase']]
                sourceDatabase = cluster[cluster_id]["databases"][cluster_json['sourceDatabase']]
                for collectionName in cluster[cluster_id]["databases"][cluster_json['sourceDatabase']]["collections"].keys():
                    tempColName = "migrated-"+collectionName+"@"+cluster_json['sourceDatabase']
                    destinationDatabase["collections"][tempColName] = {}
                    destinationDatabase["collections"][tempColName]['_uuid'] = shortuuid.uuid()
                    destinationDatabase["collections"][tempColName]['collection_name'] = tempColName
                    destinationDatabase["collections"][tempColName]['schema'] =  sourceDatabase["collections"][collectionName].get('schema')
                    destinationDatabase["collections"][tempColName]['data'] =  sourceDatabase["collections"][collectionName].get('data')
                    write_json(clusters, file)
                sourceDatabase["collections"] = {}
                write_json(clusters, file)
                return jsonify({
                            "status_code" : '1',
                            "response" : "Collections migrated."
                        })
            else:
                return jsonify({
                    "status_code" : '0',
                    "response" : "Destination Database doesn't exist."
                })
        else:
            return jsonify({
                    "status_code" : '0',
                    "response" : "Source Database doesn't exist."
                })


# MIGRATE COLLECTION
def migrate_collection(cluster_id , payload):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = payload
        if cluster_json['collection'] == "migrate-all":
            return migrate_all_collections(cluster_id, payload)
        if cluster_json['sourceDatabase'] in cluster[cluster_id]["databases"].keys():
            if cluster_json['destinationDatabase'] in cluster[cluster_id]["databases"].keys():
                sourceDatabase = cluster[cluster_id]["databases"][cluster_json['sourceDatabase']]
                destinationDatabase = cluster[cluster_id]["databases"][cluster_json['destinationDatabase']]
                if cluster_json['collection'] in cluster[cluster_id]["databases"][cluster_json['sourceDatabase']]["collections"].keys():
                    if cluster_json['collection'] in cluster[cluster_id]["databases"][cluster_json["destinationDatabase"]]["collections"].keys():
                        if cluster_json['actionIfExists'] == "replace":
                            destinationDatabase["collections"].pop(cluster_json["collection"])
                            destinationDatabase["collections"][cluster_json['collection']] = {}
                            destinationDatabase["collections"][cluster_json['collection']]['_uuid'] = shortuuid.uuid()
                            destinationDatabase["collections"][cluster_json['collection']]['collection_name'] = sourceDatabase["collections"][cluster_json['collection']].get('collection_name')
                            destinationDatabase["collections"][cluster_json['collection']]['schema'] = sourceDatabase["collections"][cluster_json['collection']].get('schema')
                            destinationDatabase["collections"][cluster_json['collection']]['data'] = sourceDatabase["collections"][cluster_json['collection']].get('data')
                            sourceDatabase["collections"].pop(cluster_json["collection"])
                            write_json(clusters, file)
                            return jsonify({
                                        "status_code" : '1',
                                        "response" : "Collection Moved!"
                                    })
                        elif cluster_json['actionIfExists'] == "rename":
                            destinationDatabase["collections"][cluster_json['newName']] = {}
                            destinationDatabase["collections"][cluster_json['newName']]['_uuid'] = shortuuid.uuid()
                            destinationDatabase["collections"][cluster_json['newName']]['collection_name'] = cluster_json['newName']
                            destinationDatabase["collections"][cluster_json['newName']]['schema'] = sourceDatabase["collections"][cluster_json['collection']].get('schema')
                            destinationDatabase["collections"][cluster_json['newName']]['data'] = sourceDatabase["collections"][cluster_json['collection']].get('data')
                            sourceDatabase["collections"].pop(cluster_json["collection"])
                            write_json(clusters, file)
                            return jsonify({
                                        "status_code" : '1',
                                        "response" : "Collection Moved!"
                                    })
                        elif  cluster_json['actionIfExists'] == "abort":
                            return jsonify({
                                        "status_code" : '0',
                                        "response" : "Process aborted as collection with same name exists in destination database."
                                    })
                    else:
                        destinationDatabase["collections"][cluster_json['collection']] = {}
                        destinationDatabase["collections"][cluster_json['collection']]['_uuid'] = shortuuid.uuid()
                        destinationDatabase["collections"][cluster_json['collection']]['collection_name'] = sourceDatabase["collections"][cluster_json['collection']].get('collection_name')
                        destinationDatabase["collections"][cluster_json['collection']]['schema'] = sourceDatabase["collections"][cluster_json['collection']].get('schema')
                        destinationDatabase["collections"][cluster_json['collection']]['data'] = sourceDatabase["collections"][cluster_json['collection']].get('data')
                        sourceDatabase["collections"].pop(cluster_json["collection"])
                        write_json(clusters, file)
                        return jsonify({
                                    "status_code" : '1',
                                    "response" : "Collection Moved!"
                                })
                else:
                    return jsonify({
                        "status_code" : '0',
                        "response" : "Collection doesn't exist."
                    })
            else:
                return jsonify({
                    "status_code" : '0',
                    "response" : "Destination Database doesn't exist."
                }) 
        else:
            return jsonify({
                    "status_code" : '0',
                    "response" : "Source Database doesn't exist."
                })


# DELETE DATABASE
def delete_database(cluster_id , payload):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = payload
        if cluster_json['targetDatabase'] in cluster[cluster_id]["databases"].keys():
            if cluster_json['migrateCollections'] == 'true':
                destinationDatabase = cluster[cluster_id]["databases"][cluster_json['destinationDatabase']]
                sourceDatabase = cluster[cluster_id]["databases"][cluster_json['targetDatabase']]
                if cluster_json['destinationDatabase'] in cluster[cluster_id]["databases"].keys():
                    if cluster_json['ifCollectionExists'] == "replace":
                        for collectionName in cluster[cluster_id]["databases"][cluster_json['targetDatabase']]["collections"].keys():
                            if collectionName in cluster[cluster_id]["databases"][cluster_json["destinationDatabase"]]["collections"].keys():
                                destinationDatabase["collections"].pop(collectionName)
                                destinationDatabase["collections"][collectionName] = {}
                                destinationDatabase["collections"][collectionName]['_uuid'] = shortuuid.uuid()
                                destinationDatabase["collections"][collectionName]['collection_name'] = sourceDatabase["collections"][collectionName].get('collection_name')
                                destinationDatabase["collections"][collectionName]['schema'] = sourceDatabase["collections"][collectionName].get('schema')
                                destinationDatabase["collections"][collectionName]['data'] = sourceDatabase["collections"][collectionName].get('data')
                            else:
                                destinationDatabase["collections"][collectionName] = {}
                                destinationDatabase["collections"][collectionName]['_uuid'] = shortuuid.uuid()
                                destinationDatabase["collections"][collectionName]['collection_name'] = sourceDatabase["collections"][collectionName].get('collection_name')
                                destinationDatabase["collections"][collectionName]['schema'] = sourceDatabase["collections"][collectionName].get('schema')
                                destinationDatabase["collections"][collectionName]['data'] = sourceDatabase["collections"][collectionName].get('data')
                            write_json(clusters, file)
                        cluster[cluster_id]["databases"].pop(cluster_json['targetDatabase'])
                        write_json(clusters, file)
                        return jsonify({
                                    "status_code" : '1',
                                    "response" : "Database deleted after migrating collections"
                                })
                    elif cluster_json['ifCollectionExists'] == "rename":
                        for collectionName in cluster[cluster_id]["databases"][cluster_json['targetDatabase']]["collections"].keys():
                            if collectionName in cluster[cluster_id]["databases"][cluster_json["destinationDatabase"]]["collections"].keys():
                                tempColName = "cloned-"+collectionName+"@"+cluster_json['targetDatabase']
                                destinationDatabase["collections"][tempColName] = {}
                                destinationDatabase["collections"][tempColName]['_uuid'] = shortuuid.uuid()
                                destinationDatabase["collections"][tempColName]['collection_name'] = tempColName
                                destinationDatabase["collections"][tempColName]['schema'] =  sourceDatabase["collections"][collectionName].get('schema')
                                destinationDatabase["collections"][tempColName]['data'] =  sourceDatabase["collections"][collectionName].get('data')
                            else:
                                destinationDatabase["collections"][collectionName] = {}
                                destinationDatabase["collections"][collectionName]['_uuid'] = shortuuid.uuid()
                                destinationDatabase["collections"][collectionName]['collection_name'] = sourceDatabase["collections"][collectionName].get('collection_name')
                                destinationDatabase["collections"][collectionName]['schema'] = sourceDatabase["collections"][collectionName].get('schema')
                                destinationDatabase["collections"][collectionName]['data'] = sourceDatabase["collections"][collectionName].get('data')
                            write_json(clusters, file)
                        cluster[cluster_id]["databases"].pop(cluster_json['targetDatabase'])
                        write_json(clusters, file)
                        return jsonify({
                                    "status_code" : '1',
                                    "response" : "Database deleted after migrating collections"
                                })
                    elif cluster_json['ifCollectionExists'] == "skip":
                        for collectionName in cluster[cluster_id]["databases"][cluster_json['targetDatabase']]["collections"].keys():
                            if collectionName in cluster[cluster_id]["databases"][cluster_json["destinationDatabase"]]["collections"].keys():
                                pass
                            else:
                                destinationDatabase["collections"][collectionName] = {}
                                destinationDatabase["collections"][collectionName]['_uuid'] = shortuuid.uuid()
                                destinationDatabase["collections"][collectionName]['collection_name'] = sourceDatabase["collections"][collectionName].get('collection_name')
                                destinationDatabase["collections"][collectionName]['schema'] = sourceDatabase["collections"][collectionName].get('schema')
                                destinationDatabase["collections"][collectionName]['data'] = sourceDatabase["collections"][collectionName].get('data')
                            write_json(clusters, file)
                        cluster[cluster_id]["databases"].pop(cluster_json['targetDatabase'])
                        write_json(clusters, file)
                        return jsonify({
                                    "status_code" : '1',
                                    "response" : "Database deleted after migrating collections"
                                }) 
                else:
                    return jsonify({
                        "status_code" : '0',
                        "response" : "Destination Database doesn't exist"
                    })
            elif cluster_json['migrateCollections'] == 'false':
                cluster[cluster_id]["databases"].pop(cluster_json['targetDatabase'])
                write_json(clusters, file)
                return jsonify({
                    "status_code" : '1',
                    "response" : "Database deleted without migration."
                })
        else:
            return jsonify({
                    "status_code" : '0',
                    "response" : "Database doesn't exist."
                })


# ADD DATA TO COLLECTION
def add_data(cluster_id , payload):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = payload
        if cluster_json['database'] in cluster[cluster_id]["databases"].keys():
            if cluster_json['collection'] in cluster[cluster_id]["databases"][cluster_json['database']]["collections"].keys():
                schema = cluster[cluster_id]["databases"][cluster_json['database']]["collections"][cluster_json['collection']]['schema']
                _id = shortuuid.uuid()
                data= cluster_json['data']
                data.update({'_id': _id})
                if set(data.keys()) == set(schema.keys()):
                    temp = cluster[cluster_id]['databases'][cluster_json['database']]['collections'][cluster_json['collection']]['data']
                    temp[_id] = data
                    write_json(clusters, file)
                    return jsonify({
                            "status_code" : '1',
                            "response" :"Data Entered"
                        })
                else:
                    return jsonify({
                            "status_code" : '0',
                            "response" : 'Collection Schema and Data Schema does not match'
                        })
            else:
                return jsonify({
                        "status_code" : '0',
                        "response" : "Collection doesn't exit."
                    })
        else:
            return jsonify({
                    "status_code" : '0',
                    "response" :"Database doesn't exist."
                })


# UPDATE DATA
def update_data(cluster_id , payload ):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = payload
        marker = cluster_json['marker'].split(" : ")
        marker_key = marker[0]
        marker_value = marker[1]
        if cluster_json['database'] in cluster[cluster_id]["databases"].keys():
            if cluster_json['collection'] in cluster[cluster_id]["databases"][cluster_json['database']]["collections"].keys():
                schema = cluster[cluster_id]["databases"][cluster_json['database']]["collections"][cluster_json['collection']]['schema']
                collection = cluster[cluster_id]["databases"][cluster_json['database']]["collections"][cluster_json['collection']]
                if marker_key in collection['schema'].keys():
                    for _id in collection['data'].keys():
                        if collection['data'][_id][marker_key] == marker_value:
                            collection['data'][_id].update(cluster_json['data'])
                            write_json(clusters, file)
                            return jsonify({
                                    "status_code" : '1',
                                    "response" :'Data Updated!'
                                })
                    return jsonify({
                            "status_code" : '0',
                            "response" :'Data Not Found' 
                        })
                else:
                    return jsonify({
                            "status_code" : '0',
                            "response" :'Marker Invalid'
                        })
            else:
                return jsonify({
                        "status_code" : '0',
                        "response" :"Collection doesn't exit."
                    })
        else:
            return jsonify({
                    "status_code" : '0',
                    "response" :"Database doesn't exist."
                })


# DELETE DATA
def delete_data(cluster_id , payload ):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = payload
        marker = cluster_json['marker'].split(" : ")
        marker_key = marker[0]
        marker_value = marker[1]
        if cluster_json['database'] in cluster[cluster_id]["databases"].keys():
            if cluster_json['collection'] in cluster[cluster_id]["databases"][cluster_json['database']]["collections"].keys():
                schema = cluster[cluster_id]["databases"][cluster_json['database']]["collections"][cluster_json['collection']]['schema']
                collection = cluster[cluster_id]["databases"][cluster_json['database']]["collections"][cluster_json['collection']]
                if marker_key in collection['schema'].keys():
                    for _id in collection['data'].keys():
                        if collection['data'][_id][marker_key] == marker_value:
                            collection['data'].pop(_id)
                            write_json(clusters, file)
                            return jsonify({
                                    "status_code" : '1',
                                    "response" : 'Data Deleted!'
                                })
                            break
                    return jsonify({
                            "status_code" : '0',
                            "response" : 'Data Not Found' 
                        })
                else:
                    return jsonify({
                            "status_code" : '0',
                            "response" : 'Marker Invalid'
                        })
            else:
                return jsonify({
                        "status_code" : '0',
                        "response" :"Collection doesn't exit."
                    })
        else:
            return jsonify({
                    "status_code" : '0',
                    "response" :"Database doesn't exist."
                })

# GET COMPLETE CLUSTER
def get_cluster(cluster_id , payload ):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = payload
        return jsonify({ "status_code" : "1" , "cluster" : cluster })

            
# GET DATABASES
def get_databases(cluster_id , payload ):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = payload
        if cluster_json['database'] == 'all' :
            result = []
            for i in cluster[cluster_id]["databases"].keys():  
                result.append({
                    "status_code" : "1",
                    "database_name" : cluster[cluster_id]["databases"][i]['db_name'],
                    "_uuid" : i,
                    "collections_count": len(cluster[cluster_id]["databases"][i]['collections'])
                })
            return jsonify({ "response" : result})
        else:
            if cluster_json['database'] in cluster[cluster_id]["databases"].keys():
                return jsonify({
                    "status_code" : "1",
                    "database_name" : cluster[cluster_id]["databases"][cluster_json['database']]['db_name'],
                    "_uuid" : cluster_json['database'],
                    "collections_count": len(cluster[cluster_id]["databases"][cluster_json['database']]['collections'])
                })
            else:
                return jsonify({
                    "status_code" : '0',
                    "response" :"Database doesn't exist."
                })


# GET COLLECTIONS
def get_collections(cluster_id , payload ):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = payload
        if cluster_json['database'] in cluster[cluster_id]["databases"].keys():
            if cluster_json['collection'] == 'all' :
                result = []
                for i in cluster[cluster_id]["databases"][cluster_json['database']]['collections'].keys():  
                    result.append({
                        "collection_name" : cluster[cluster_id]["databases"][cluster_json['database']]['collections'][i]['collection_name'],
                        "_uuid" : i,
                        "data_count": len(cluster[cluster_id]["databases"][cluster_json['database']]['collections'][i]['data'])
                    })
                return jsonify({"status_code" : "1", "response" : result})
            else:
                if cluster_json['collection'] in cluster[cluster_id]["databases"][cluster_json['database']]['collections'].keys():
                    return jsonify({
                        "collection_name" : cluster[cluster_id]["databases"][cluster_json['database']]['collections'][cluster_json['collection']]['collection_name'],
                        "_uuid" : cluster_json['collection'],
                        "data_count": len(cluster[cluster_id]["databases"][cluster_json['database']]['collections'][cluster_json['collection']]['data'])
                    })
                else:
                    return jsonify({
                        "status_code" : '0',
                        "response" :"Collection doesn't exit."
                    })
        else:
            return jsonify({
                    "status_code" : '0',
                    "response" :"Database doesn't exist."
                })


# GET DATA
def get_data(cluster_id , payload ):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = payload
        if cluster_json['database'] in cluster[cluster_id]["databases"].keys():
            if cluster_json['collection'] in cluster[cluster_id]["databases"][cluster_json['database']]["collections"].keys():
                data = cluster[cluster_id]["databases"][cluster_json['database']]["collections"][cluster_json['collection']]['data']
                schema = cluster[cluster_id]["databases"][cluster_json['database']]["collections"][cluster_json['collection']]['schema']
                return jsonify({
                    "status_code" : "1",
                    "data" : data,
                    "schema" : schema
                })
            else:
                return jsonify({
                        "status_code" : '0',
                        "response" :"Collection doesn't exit."
                    })
        else:
            return jsonify({
                    "status_code" : '0',
                    "response" :"Database doesn't exist."
                })
           
# CONNECT - SINGLE ENDPOINT
@app.route('/connect', methods = ['GET' , 'POST'])
def connect():
    cluster_id = request.args.get('cluster')
    token = request.args.get('token')
    file = "./clusters/"+cluster_id+".json"
    if path.exists(file):
        with open(file) as db:
            body = json.load(db)
            cluster = body['clusters']
            if token in cluster[cluster_id]['access_tokens']:
                action_json = request.get_json(force = True)
                payload = action_json['payload']
                if action_json['action'].lower() == "create-database" :
                    return create_database(cluster_id, payload)
                elif action_json['action'].lower() == "create-collection" :
                    return create_collection(cluster_id, payload)
                elif action_json['action'].lower() == "add-data" :
                    return add_data(cluster_id, payload)
                elif action_json['action'].lower() == "update-data" :
                    return update_data(cluster_id, payload)
                elif action_json['action'].lower() == "delete-data" :
                    return delete_data(cluster_id, payload)
                elif action_json['action'].lower() == "get-cluster" :
                    return get_cluster(cluster_id, payload)
                elif action_json['action'].lower() == "get-collection" :
                    return get_collections(cluster_id, payload)
                elif action_json['action'].lower() == "get-database" :
                    return get_databases(cluster_id, payload)
                elif action_json['action'].lower() == "get-data" :
                    return get_data(cluster_id, payload)
                elif action_json["action"].lower() == "delete-collection":
                    return delete_collection(cluster_id, payload)
                elif action_json["action"].lower() == "clone-collection":
                    return clone_collection(cluster_id, payload)
                elif action_json["action"].lower() == "migrate-collection":
                    return migrate_collection(cluster_id, payload)
                elif action_json["action"].lower() == "delete-database":
                    return delete_database(cluster_id, payload)
            else:
                return  jsonify({ "status_code" : "0", "response" : 'Wrong Access token'})
    else:
        return jsonify({"status_code" : "0",  "response" : 'Cluster does not exist.' }) 
        
        
if __name__ == '__main__':
    app.run(debug=True)