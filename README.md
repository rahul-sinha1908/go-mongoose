# Go Mongoose Driver

This is the driver made over mongo db to ease many operations like populate and so with mongo db

## Things to know about this new Wrapper over MongoDB

There are certain rules that is required to be followed for perfect setup of the Mongo DB-
1. Need to specify all the struct to the mongoose init statement a the start of the application so that Its relations are setup
1. MSON tags plays a great role to make everthing work perfectly for mongoose

## MSON tags

These are some tags that would help in speeding up certain populate statements

- collection : This tag is used to specify the collection its refereing to \
Example - 
```
type XYZ struc{
    ObjRef primitive.ObjectId `json:"objRef" bson:"objRef" mson:"collection='refCollection'"`
}
```
- unique : This Keyword should be used when the key is unique
```
type XYZ struc{
    ObjRef primitive.ObjectId `json:"objRef" bson:"objRef" mson:"unique"`
}
```
- cunique : This Keyword is used When there is a combined Unique key in the platform
```
type XYZ struc{
    ObjRef primitive.ObjectId `json:"objRef" bson:"objRef" mson:"cunique"`
    Title string `json:"title" bson:"title" mson:"cunique"`
}
```
