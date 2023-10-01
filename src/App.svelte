<script>
	import { onMount } from 'svelte';
	import Header from './Header.svelte';
	import SearchBar from './SearchBar.svelte';
	import SearchResult from './searchResult.svelte'; 
	import Footer from './Footer.svelte';
	import SERP from './SERP.svelte';
	import Chatbox from './Chatbot.svelte';
	import SchoolButton from './SchoolButton.svelte';
	import CollegeButton from './CollegeButton.svelte';
	import UniversityButton from './UniversityButton.svelte';
	import CoursesButton from './CoursesButton.svelte';
	let searchQuery = '';
	console.log(searchQuery)
	let searchResults = [];
	let filteredSearchResults = [];


	let user = {
    name: '',
    email: '',
    number: '',
    password: '',
  };

	async function getData(searchQuery) {
    searchQuery = searchQuery.trim();

    if (searchQuery === '') {
      searchResults = [];
      return;
    }
    const specialCharactersRegex = /[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]/;
    if (specialCharactersRegex.test(searchQuery)) {
      searchResults = [];
      return;
    }

    try {
      const response = await fetch(`http://localhost:8081/api/search?query=${searchQuery}`);
      if (!response.ok) {
        throw new Error('Failed to fetch search results');
      }

      const data = await response.json();
      searchResults = data;
      console.log(data);
    } catch (error) {
      console.error(error);
      searchResults = []; 
    }
  }

	
	function handleSearchs(event){
		searchQuery = event.detail.source
		getData(searchQuery)
		console.log(searchQuery)
	}

	const tokenizeQuery = () => {
		searchQuery = searchQuery.toLowerCase(); 
		const tokens = searchQuery.split(/\s+/); 
		return tokens;
	};

	onMount(() => {
		tokenizeQuery();
	});


	function handleMessage(event){
		console.log(event)
	}
</script>

<main>
	<Header />
	<div class="app-container"></div>
	{#if searchQuery === ""}
	<div style="margin-bottom: 90px;"></div>
	<img src="C:/Users/Admin/OneDrive/Desktop/Correct Code/Trial 2/svelte/public/bird.png" alt="Logo" class="logo" />
	<img src="C:/Users/Admin/OneDrive/Desktop/Correct Code/Trial 2/svelte/public/descr.png" alt="Description" class="Description" />	
	<SearchBar bind:searchQuery={searchQuery} on:cusEvent={handleSearchs}/>
	<div style="display: flex; margin-top: 10px;">
		<SchoolButton label="Schools" />
		<div style="width: 20px;"></div> <!-- Space between buttons -->
		<CollegeButton label="Colleges" />
		<div style="width: 20px;"></div> <!-- Space between buttons -->
		<UniversityButton label="Universities" />
		<div style="width: 20px;"></div> <!-- Space between buttons -->
		<CoursesButton label="Courses" />
	  </div>
	<SearchResult {searchResults} {filteredSearchResults} /> 
	{:else}
	<SERP {searchResults} />
	{/if}
	<Chatbox />
	<Footer />
	
</main>
<style>
	body, html {
    margin: 0;
  }

  main {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 5px; /* Adjust the padding as desired */
  }

  .logo {
    width: 100px;
    height: 100px;
    margin-bottom: 4px;
    padding: 0; /* Adjust the padding as desired */
  }
  .Description{
	width: 400px;
    height: 50px;
    margin-bottom: 4px;
    padding: 0; /* Adjust the padding as desired */
  }

  
</style>
