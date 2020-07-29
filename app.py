import json
import shortuuid
import secrets
from flask import Flask, request, jsonify, send_from_directory

app = Flask(__name__)

def write_json(data, filename="db.json"):
    with open(filename, 'w') as f:
        json.dump(data, f, indent=4)

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
                        "db_name": db_json['db_name'],
                        "collections": []
                    }
                    temp = i['databases']
                    temp.append(db_data)
                    write_json(clusters)
                    return jsonify({ 
                        "response" : "Database Created with ID "+database_id
                        })
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
                                "collection_name": cluster_json['collection_name'],
                                "schema": schema,
                                "data": []
                            }
                            temp = j['collections']
                            temp.append(collection_data)
                            write_json(clusters)
                            return jsonify({ 
                                "response" : "Collection Created with ID "+collection_id
                                })
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
                                        return jsonify({
                                            "response" : "Data added"
                                            })
                                    else:
                                        return 'Collection Schema and Data Schema does not match'
                            return 'No Collection found with this ID.'
                    return 'No Database found with this ID.'
                else:
                    return 'Wrong Access token'
        return 'Cluster does not exist.'










'''# READ
@app.route('/<api_key>/projects/view', methods = ['GET'])
def view_projects(api_key):
    with open('keys.json') as key_json:
        keys = json.load(key_json)
        if api_key in keys['api_keys']:
            with open('projects.json') as json_file:
                data = json.load(json_file)
                projects = data['projects']
            return jsonify(projects)
        else:
            return 'Wrong API Key'

# UPDATE
@app.route('/<api_key>/projects/update', methods = ['POST'])
def update_projects(api_key):
    with open('keys.json') as key_json:
        keys = json.load(key_json)
        data_json = request.get_json(force = True)
        if api_key in keys['api_keys']:
            _uuid = data_json['_uuid']
            with open('projects.json') as json_file:
                data = json.load(json_file)
                temp = data['projects']
                if not any(d['_uuid'] == _uuid for d in temp):
                    return 'Data does not exist'
                else:
                    for i in range(len(temp)): 
                        if temp[i]['_uuid'] == _uuid: 
                            temp[i].update(data_json)
                            break
                    write_json(data)
                    return 'Updated!'
        else:
            return 'Wrong API Key'

# DELETE
@app.route('/<api_key>/projects/delete', methods = ['POST'])
def delete_projects(api_key):
    with open('keys.json') as key_json:
        keys = json.load(key_json)
        data_json = request.get_json(force = True)
        if api_key in keys['api_keys']:
            _uuid = data_json['_uuid']
            with open('projects.json') as json_file:
                data = json.load(json_file)
                temp = data['projects']

                if not any(d['_uuid'] == _uuid for d in temp):
                    return 'Data does not exist'
                else:
                    for i in range(len(temp)): 
                        if temp[i]['_uuid'] == _uuid: 
                            del temp[i] 
                            break
                    write_json(data)
                    return 'Deleted!'
        else:
            return 'Wrong API Key'
'''