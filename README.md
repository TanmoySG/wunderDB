# wunderDB
wunderDB is a JSON-based micro Document DB, inspired by MongoDB.

It uses **Cluster->Database->Collection->Data** hierarchy to store data.

* ***Cluster*** is an group of databases pertaining to one user.
* ***Databases*** are group of collection.
* ***Collections***, like Tables in databases, store series of data pertaining to a single domain.

version ***0.1 Beta***

developed by ***[Tanmoy Sen Gupta](https://www.tanmoysg.com)***

Get started by downloading the **app.py and db.json** files.
* Create a **Python Virtual Environment**.
* Install **Flask** and **shortuuid** using pip.
* Run '***flask run***' and your DB is ready to use!

When the development server starts access the instructions by going to **localhost:5000/get-started**

To be deployed/hosted at **[www.tanmoysg.com](https://www.tanmoysg.com)**
