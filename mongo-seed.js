conn = new Mongo();
db = conn.getDB("pets");

// add user
if (db.users.count() == 0) {
	var user = {name:"bob", pwd:"dylan"};
	insert(user, "users")
}

// add some pets
if (db.pets.count() == 0) {
	var pets = [
		{name:"bixa", kind:"cat"},
		{name:"rex", kind:"cat"}
	];
	insert(pets, "pets")
}

function insert(payload, collection) {
	db[collection].insert(payload, {}, function(){
		print("mongo-seed: " + collection + " updated")
	})
}
