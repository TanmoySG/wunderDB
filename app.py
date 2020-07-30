
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
from flask import Flask, request, jsonify, send_from_directory

app = Flask(__name__)

def write_json(data, filename="db.json"):
    with open(filename, 'w') as f:
        json.dump(data, f, indent=4)

# GETTING STARTED
@app.route('/get-started', methods = ['GET'])
def get_started():
    with open('db.json') as db:
        instructions = json.load(db)
        instruction = instructions['templates']
        return jsonify({
            "_00_message" : "Welcome to wunderDB !",
            "_01_about"   : "wunderDB is a JSON-based micro Document DB inspired by MongoDB.",
            "_02_version" : "0.1 Beta",
            "_03_creator" : "Tanmoy Sen Gupta",
            "_04_instructions" : instruction
        })
        
# CREATE CLUSTER
@app.route('/create/cluster', methods = ['POST'])
def create_cluster():
    user_data = request.get_json(force = True)
    tokens = []
    cluster_id =shortuuid.uuid()
    for i in range(3):
        tokens.append(secrets.token_hex(16))
    cluster_data = {
        "_cluster_id": cluster_id,
        "username": user_data['username'],
        "password": user_data['password'],
        "access_tokens": tokens,
        "databases":[]
    }
    with open('db.json') as json_file:
        data = json.load(json_file)
        temp = data['clusters']
        temp.append(cluster_data)
        write_json(data)
    return jsonify({ 
        "response" : "Cluster Created with ID "+cluster_id , 
        "access_tokens" : tokens
        })


# CREATE DATABASE
@app.route('/<cluster_id>/<access_token>/create/database', methods = ['POST'])
def create_database(cluster_id , access_token):
    with open('db.json') as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        for i in cluster:
            if i['_cluster_id'] == cluster_id :
                if access_token in i['access_tokens']:
                    db_json = request.get_json(force = True)
                    database_id =shortuuid.uuid()
                    db_data={
                        "_uuid": database_id,
                        "db_name": db_json['name'],
                        "collections": []
                    }
                    temp = i['databases']
                    temp.append(db_data)
                    write_json(clusters)
                    return "Database Created with ID "+database_id
                else:
                    return 'Wrong Access token'
        return 'Cluster does not exist.'

# CREATE COLLECTION
@app.route('/<cluster_id>/<access_token>/create/collection', methods = ['POST'])
def create_collection(cluster_id , access_token):
    with open('db.json') as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        for i in cluster:
            if i['_cluster_id'] == cluster_id :
                if access_token in i['access_tokens']:
                    cluster_json = request.get_json(force = True)
                    for j in i['databases']:
                        if j['_uuid'] == cluster_json['db_id'] :
                            collection_id =shortuuid.uuid()
                            schema = cluster_json['schema']
                            schema.update({"_id" : "ID"})
                            collection_data={
                                "_uuid": collection_id,
                                "collection_name": cluster_json['name'],
                                "schema": schema,
                                "data": []
                            }
                            temp = j['collections']
                            temp.append(collection_data)
                            write_json(clusters)
                            return "Collection Created with ID "+collection_id
                    return 'No Database found with this id.'
                else:
                    return 'Wrong Access token'
        return 'Cluster does not exist.'

# ADD DATA TO COLLECTION
@app.route('/<cluster_id>/<access_token>/add/data', methods = ['POST'])
def add_data(cluster_id , access_token):
    with open('db.json') as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        for i in cluster:
            if i['_cluster_id'] == cluster_id :
                if access_token in i['access_tokens']:
                    collection_json = request.get_json(force = True)
                    for j in i['databases']:
                        if j['_uuid'] == collection_json['db_id'] :
                            for k in j['collections']:
                                if k['_uuid'] == collection_json['collection_id']:
                                    data_id = shortuuid.uuid()
                                    data= collection_json['data']
                                    data.update({'_id': data_id})
                                    if set(data.keys()) == set(k['schema'].keys()):
                                        temp = k['data']
                                        temp.append(data)
                                        write_json(clusters)
                                        return "Data added"
                                    else:
                                        return 'Collection Schema and Data Schema does not match'
                            return 'No Collection found with this ID.'
                    return 'No Database found with this ID.'
                else:
                    return 'Wrong Access token'
        return 'Cluster does not exist.'

# UPDATE DATA
@app.route('/<cluster_id>/<access_token>/update/data', methods = ['POST'])
def update_data(cluster_id , access_token ):
    with open('db.json') as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        for i in cluster:
            if i['_cluster_id'] == cluster_id :
                if access_token in i['access_tokens']:
                    update_json = request.get_json(force = True)
                    marker = update_json['marker'].split(" : ")
                    marker_key = marker[0]
                    marker_value = marker[1]
                    for j in i['databases']:
                        if j['_uuid'] == update_json['db_id'] :
                            for k in j['collections']:
                                if k['_uuid'] == update_json['collection_id']:
                                    if marker_key in k['schema'].keys():
                                        for l in k['data']:
                                            if l[marker_key] == marker_value:
                                                l.update(update_json['data'])
                                                break
                                        write_json(clusters)
                                        return 'Data Updated!'
                                    else:
                                        return 'Marker Invalid'    
                            return 'No Collection found with this ID.'
                    return 'No Database found with this ID.'
                else:
                    return 'Wrong Access token'
        return 'Cluster does not exist.'

# DELETE DATA
@app.route('/<cluster_id>/<access_token>/delete/data', methods = ['POST'])
def delete_data(cluster_id , access_token ):
    with open('db.json') as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        for i in cluster:
            if i['_cluster_id'] == cluster_id :
                if access_token in i['access_tokens']:
                    delete_json = request.get_json(force = True)
                    marker = delete_json['marker'].split(" : ")
                    marker_key = marker[0]
                    marker_value = marker[1]
                    for j in i['databases']:
                        if j['_uuid'] == delete_json['db_id'] :
                            for k in j['collections']:
                                if k['_uuid'] == delete_json['collection_id']:
                                    if marker_key in k['schema'].keys():
                                        for l in range(len(k['data'])):
                                            if k['data'][l][marker_key] == marker_value:
                                                del k['data'][l]
                                                write_json(clusters)
                                                return 'Data Deleted!'
                                                break
                                        return 'Data not found.'
                                    else:
                                        return 'Marker Invalid'    
                            return 'No Collection found with this ID.'
                    return 'No Database found with this ID.'
                else:
                    return 'Wrong Access token'
        return 'Cluster does not exist.'

# READ DATA
@app.route('/<cluster_id>/<access_token>/view/data', methods = ['GET'])
def get_data(cluster_id , access_token ):
    with open('db.json') as db:
        clusters = json.load(db)
        cluster = clusters['clusters']
        for i in cluster:
            if i['_cluster_id'] == cluster_id :
                if access_token in i['access_tokens']:
                    query_json = request.get_json(force = True)
                    for j in i['databases']:
                        if j['_uuid'] == query_json['db_id'] :
                            for k in j['collections']:
                                if k['_uuid'] == query_json['collection_id'] :
                                    return jsonify({ 
                                        '_name': k['collection_name'],
                                        'schema' : k['schema'] ,
                                        'data': k['data'] 
                                        })
                            return 'No Collection found with this ID.'
                    return 'No Database found with this ID.'
                else:
                    return 'Wrong Access token'
        return 'Cluster does not exist.'
