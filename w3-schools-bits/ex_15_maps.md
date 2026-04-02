## Go Maps

Maps are used to store data values in key:value pairs.

Each element in a map is a key:value pair.

A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

Go has multiple ways for creating maps.
Create Maps Using var and :=
Syntax
var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}



## Note: The order of the map elements defined in the code is different from the way that they are stored. The data are stored in a way to have efficient data retrieval from the map.

Create Maps Using the make() Function:
Syntax
var a = make(map[KeyType]ValueType)
b := make(map[KeyType]ValueType)

Create an Empty Map

There are two ways to create an empty map. One is by using the make()function and the other is by using the following syntax.
Syntax
var a map[KeyType]ValueType

## Note: The make()function is the right way to create an empty map. If you make an empty map in a different way and write to it, it will causes a runtime panic.

## Allowed Key Types

The map key can be of any data type for which the equality operator (==) is defined. These include:

    Booleans
    Numbers
    Strings
    Arrays
    Pointers
    Structs
    Interfaces (as long as the dynamic type supports equality)

## Invalid key types are:

    Slices
    Maps
    Functions

These types are invalid because the equality operator (==) is not defined for them.

## Allowed Value Types

The map values can be any type.
Access Map Elements

You can access map elements by:
Syntax
value = map_name[key]

## Update and Add Map Elements

Updating or adding an elements are done by:
Syntax
map_name[key] = value

Remove Element from Map

## Removing elements is done using the delete() function.
Syntax
delete(map_name, key)

## Check For Specific Elements in a Map

You can check if a certain key exists in a map using:
Syntax
val, ok :=map_name[key]

If you only want to check the existence of a certain key, you can use the blank identifier (_) in place of val.

## Maps Are References

Maps are references to hash tables.

If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.

## Iterate Over Maps

You can use range to iterate over maps.

## Iterate Over Maps in a Specific Order

Maps are unordered data structures. If you need to iterate over a map in a specific order, you must have a separate data structure that specifies that order. 

