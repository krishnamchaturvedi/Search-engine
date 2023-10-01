import { openDB } from 'idb';

// Database configuration
const DB_NAME = 'my_database';
const DB_VERSION = 1;
const OBJECT_STORE_NAME = 'my_object_store';

// Function to open the database and create the object store
export async function openDatabase() {
  return openDB(DB_NAME, DB_VERSION, {
    upgrade(db) {
      if (!db.objectStoreNames.contains(OBJECT_STORE_NAME)) {
        db.createObjectStore(OBJECT_STORE_NAME, { keyPath: 'title', autoIncrement: true });
      }
    },
  });
}

// Function to perform CRUD operations

// Create a new item in the object store
export async function addItem(item) {
  try {
    const db = await openDatabase();
    const tx = db.transaction(OBJECT_STORE_NAME, 'readwrite');
    const store = tx.objectStore(OBJECT_STORE_NAME);
    await store.add(item);
    await tx.complete; // Commit the transaction
  } catch (error) {
    console.error('Error adding item:', error);
  }
}
// Read all items from the object store
export async function getAllItems() {
  const db = await openDatabase();
  const tx = db.transaction(OBJECT_STORE_NAME, 'readonly');
  const store = tx.objectStore(OBJECT_STORE_NAME);
  return store.getAll();
}

// Update an existing item in the object store
export async function updateItem(item) {
  const db = await openDatabase();
  const tx = db.transaction(OBJECT_STORE_NAME, 'readwrite');
  const store = tx.objectStore(OBJECT_STORE_NAME);
  await store.put(item);
}

// Delete an item from the object store
export async function deleteItem(itemId) {
  const db = await openDatabase();
  const tx = db.transaction(OBJECT_STORE_NAME, 'readwrite');
  const store = tx.objectStore(OBJECT_STORE_NAME);
  await store.delete(itemId);
}
