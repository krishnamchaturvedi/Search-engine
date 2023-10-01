<!-- searchResult.svelte -->
<script>
  // Props to receive the search results from the parent component
  export let searchResults;
</script>

<div class="search-results">
  {#if searchResults && searchResults.length > 0}
    <!-- Display all search results if no filtered results -->
    {#each searchResults as result}
      <!-- Display individual search result data here -->
      <div class="search-result">
        <p class="result-title">{result.title}</p>
        <p class="result-description">{result.description}</p>
        <p class="result-url">{result.url}</p>
      </div>
    {/each}
  <!-- {:else} -->
    <!-- <p>No results found.</p> -->
  {/if}
</div>

<style>
  /* Add styles here to format the search results */
  .search-results {
    margin-top: 20px;
  }

  .search-result {
    margin-bottom: 20px;
    padding: 10px;
    border: 1px solid #ddd;
    border-radius: 4px;
  }

  .result-title {
    font-size: 18px;
    font-weight: bold;
    color: #1a0dab;
    margin-bottom: 5px;
  }

  .result-description {
    font-size: 16px;
    color: #545454;
    margin-bottom: 5px;
  }

  .result-url {
    font-size: 14px;
    color: #006621;
    text-decoration: none;
    display: block;
  }

  .result-url:hover {
    text-decoration: underline;
  }
</style>
