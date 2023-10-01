// service-worker.js

const CACHE_NAME = 'my-svelte-app-cache-v1';
const urlsToCache = [
  '/', // Change this to the root URL of your app
  '/index.html', // Add any other files or paths you want to cache
  '/global.css', // Cache the main CSS file
  '/build/bundle.js', // Cache the main JavaScript bundle
  // Add more URLs to cache here, such as images, fonts, or other assets
];

self.addEventListener('install', (event) => {
  event.waitUntil(
    caches.open(CACHE_NAME)
      .then((cache) => cache.addAll(urlsToCache))
      .then(() => self.skipWaiting())
  );
});

self.addEventListener('activate', (event) => {
  event.waitUntil(
    caches.keys()
      .then((cacheNames) => {
        return Promise.all(
          cacheNames.map((cacheName) => {
            if (cacheName !== CACHE_NAME) {
              return caches.delete(cacheName);
            }
          })
        );
      })
      .then(() => self.clients.claim())
  );
});

self.addEventListener('fetch', (event) => {
  event.respondWith(
    caches.match(event.request)
      .then((response) => {
        // Cache hit - return the cached response
        if (response) {
          return response;
        }

        // If the request is not in the cache, fetch it from the network
        return fetch(event.request);
      })
  );
});
