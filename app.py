
'''
#############################  wunderDB  #############################

wunderDB is a JSON-based micro Document DB, inspired by MongoDB.
It uses Cluster -> Database -> Collection -> Data hierarchy to store data.

version 0.1 Beta
developed by Tanmoy Sen Gupta

* Get started by downloading the app.py and db.json files.
* Create a Python Virtual Environment.
* Install Flask and shortuuid using pip.
* Run 'flask run' and your DB is ready to use!

When the development server starts access the instructions by going to localhost:5000/get-started

'''


import json
import shortuuid
import secrets
import random
import hashlib
from flask import Flask, request, jsonify, send_from_directory

app = Flask(__name__)

def write_json(data, filename):
    with open(filename, 'w') as f:
        json.dump(data, f, indent=4)

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
@app.route('/register', methods = ['POST'])
def register():
    user_data = request.get_json(force = True)
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
                "response" : "Cluster Created",
                "cluster_id" : cluster_id , 
                "access_tokens" : tokens
                })
        else:
            return jsonify({ "response" : "User already exists." })

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
                    "response" : "Logged-in",
                    "name" : user[user_data["username"]]['name'],
                    "cluster_id" : user[user_data["username"]]['_cluster_id'],
                    "access_token" : random.choice(user[user_data["username"]]['access_tokens'])
                })
            else:
                return jsonify({
                "response" : "Password is wrong."
            })
        else:
            return jsonify({
                "response" : "User doesn't exist"
            })

# CREATE DATABASE
@app.route('/<cluster_id>/<access_token>/create/database', methods = ['POST'])
def create_database(cluster_id , access_token):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        db_json = request.get_json(force = True)
        if cluster_id in cluster.keys():
            if access_token in cluster[cluster_id]['access_tokens']:
                database_id =shortuuid.uuid()
                db_data={
                    "_uuid": database_id,
                    "db_name": db_json['name'],
                    "collections": {}
                }
                temp = cluster[cluster_id]['databases']
                temp[database_id] = db_data
                write_json(clusters, file)
                return "Database Created with ID "+database_id
            else:
                return 'Wrong Access token'
        else:
            return 'Cluster does not exist.'

# CREATE COLLECTION
@app.route('/<cluster_id>/<access_token>/create/collection', methods = ['POST'])
def create_collection(cluster_id , access_token):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = request.get_json(force = True)
        if cluster_id in cluster.keys():
            if access_token in cluster[cluster_id]['access_tokens']:
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
                    temp[collection_id] = collection_data
                    write_json(clusters, file)
                    return "Collection created with ID "+ collection_id
                else:
                    return "Database doesn't exist."
            else:
                return 'Wrong Access token'
        else:
            return 'Cluster does not exist.'

# ADD DATA TO COLLECTION
@app.route('/<cluster_id>/<access_token>/add/data', methods = ['POST'])
def add_data(cluster_id , access_token):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = request.get_json(force = True)
        if cluster_id in cluster.keys():
            if access_token in cluster[cluster_id]['access_tokens']:
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
                            return "Data Entered"
                        else:
                           return 'Collection Schema and Data Schema does not match'
                    else:
                        return "Collection doesn't exit."
                else:
                    return "Database doesn't exist."
            else:
                return 'Wrong Access token'
        else:
            return 'Cluster does not exist.'

# UPDATE DATA
@app.route('/<cluster_id>/<access_token>/update/data', methods = ['POST'])
def update_data(cluster_id , access_token ):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = request.get_json(force = True)
        marker = cluster_json['marker'].split(" : ")
        marker_key = marker[0]
        marker_value = marker[1]
        if cluster_id in cluster.keys():
            if access_token in cluster[cluster_id]['access_tokens']:
                if cluster_json['database'] in cluster[cluster_id]["databases"].keys():
                    if cluster_json['collection'] in cluster[cluster_id]["databases"][cluster_json['database']]["collections"].keys():
                        schema = cluster[cluster_id]["databases"][cluster_json['database']]["collections"][cluster_json['collection']]['schema']
                        collection = cluster[cluster_id]["databases"][cluster_json['database']]["collections"][cluster_json['collection']]
                        if marker_key in collection['schema'].keys():
                            for _id in collection['data'].keys():
                                if collection['data'][_id][marker_key] == marker_value:
                                    collection['data'][_id].update(cluster_json['data'])
                                    write_json(clusters, file)
                                    return 'Data Updated!'
                                    break
                            return 'Data Not Found' 
                        else:
                           return 'Marker Invalid'
                    else:
                        return "Collection doesn't exit."
                else:
                    return "Database doesn't exist."
            else:
                return 'Wrong Access token'
        else:
            return 'Cluster does not exist.'

# DELETE DATA
@app.route('/<cluster_id>/<access_token>/delete/data', methods = ['POST'])
def delete_data(cluster_id , access_token ):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = request.get_json(force = True)
        marker = cluster_json['marker'].split(" : ")
        marker_key = marker[0]
        marker_value = marker[1]
        if cluster_id in cluster.keys():
            if access_token in cluster[cluster_id]['access_tokens']:
                if cluster_json['database'] in cluster[cluster_id]["databases"].keys():
                    if cluster_json['collection'] in cluster[cluster_id]["databases"][cluster_json['database']]["collections"].keys():
                        schema = cluster[cluster_id]["databases"][cluster_json['database']]["collections"][cluster_json['collection']]['schema']
                        collection = cluster[cluster_id]["databases"][cluster_json['database']]["collections"][cluster_json['collection']]
                        if marker_key in collection['schema'].keys():
                            for _id in collection['data'].keys():
                                if collection['data'][_id][marker_key] == marker_value:
                                    collection['data'].pop(_id)
                                    write_json(clusters, file)
                                    return 'Data Deleted!'
                                    break
                            return 'Data Not Found' 
                        else:
                           return 'Marker Invalid'
                    else:
                        return "Collection doesn't exit."
                else:
                    return "Database doesn't exist."
            else:
                return 'Wrong Access token'
        else:
            return 'Cluster does not exist.'

# GET COMPLETE CLUSTER
@app.route('/<cluster_id>/<access_token>/get/cluster', methods = ['GET'])
def get_cluster(cluster_id , access_token ):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = request.get_json(force = True)
        if cluster_id in cluster.keys():
            if access_token in cluster[cluster_id]['access_tokens']:
                return jsonify({ "cluster" : cluster })
            else:
                return 'Wrong Access token'
        else:
            return 'Cluster does not exist.'
            
# GET DATABASES
@app.route('/<cluster_id>/<access_token>/get/databases', methods = ['GET'])
def get_databases(cluster_id , access_token ):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = request.get_json(force = True)
        if cluster_id in cluster.keys():
            if access_token in cluster[cluster_id]['access_tokens']:
                if cluster_json['database'] == 'all' :
                    result = []
                    for i in cluster[cluster_id]["databases"].keys():  
                        result.append({
                            "database_name" : cluster[cluster_id]["databases"][i]['db_name'],
                            "_uuid" : i,
                            "collections_count": len(cluster[cluster_id]["databases"][i]['collections'])
                        })
                    return jsonify({ "response" : result})
                else:
                    if cluster_json['database'] in cluster[cluster_id]["databases"].keys():
                        return jsonify({
                            "database_name" : cluster[cluster_id]["databases"][cluster_json['database']]['db_name'],
                            "_uuid" : cluster_json['database'],
                            "collections_count": len(cluster[cluster_id]["databases"][cluster_json['database']]['collections'])
                        })
                    else:
                        return 'No Database found with this ID.'
            else:
                return 'Wrong Access token'
        else:
            return 'Cluster does not exist.'

# GET COLLECTIONS
@app.route('/<cluster_id>/<access_token>/get/collections', methods = ['GET'])
def get_collections(cluster_id , access_token ):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = request.get_json(force = True)
        if cluster_id in cluster.keys():
            if access_token in cluster[cluster_id]['access_tokens']:
                if cluster_json['database'] in cluster[cluster_id]["databases"].keys():
                    if cluster_json['collection'] == 'all' :
                        result = []
                        for i in cluster[cluster_id]["databases"][cluster_json['database']]['collections'].keys():  
                            result.append({
                                "collection_name" : cluster[cluster_id]["databases"][cluster_json['database']]['collections'][i]['collection_name'],
                                "_uuid" : i,
                                "data_count": len(cluster[cluster_id]["databases"][cluster_json['database']]['collections'][i]['data'])
                            })
                        return jsonify({ "response" : result})
                    else:
                        if cluster_json['collection'] in cluster[cluster_id]["databases"][cluster_json['database']]['collections'].keys():
                            return jsonify({
                                "collection_name" : cluster[cluster_id]["databases"][cluster_json['database']]['collections'][cluster_json['collection']]['collection_name'],
                                "_uuid" : cluster_json['collection'],
                                "data_count": len(cluster[cluster_id]["databases"][cluster_json['database']]['collections'][cluster_json['collection']]['data'])
                            })
                        else:
                            return 'No Collection found with this ID.'
                else:
                    return 'No database with this ID Found.'
            else:
                return 'Wrong Access token'
        else:
            return 'Cluster does not exist.'

# GET DATA
@app.route('/<cluster_id>/<access_token>/get/data', methods = ['GET'])
def get_data(cluster_id , access_token ):
    file = "./clusters/"+cluster_id+".json"
    with open(file) as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        cluster_json = request.get_json(force = True)
        if cluster_id in cluster.keys():
            if access_token in cluster[cluster_id]['access_tokens']:
                if cluster_json['database'] in cluster[cluster_id]["databases"].keys():
                    if cluster_json['collection'] in cluster[cluster_id]["databases"][cluster_json['database']]["collections"].keys():
                        data = cluster[cluster_id]["databases"][cluster_json['database']]["collections"][cluster_json['collection']]['data']
                        return jsonify({
                            "data" : data
                        })
                    else:
                        return "Collection doesn't exit."
                else:
                    return "Database doesn't exist."
            else:
                return 'Wrong Access token'
        else:
            return 'Cluster does not exist.'
