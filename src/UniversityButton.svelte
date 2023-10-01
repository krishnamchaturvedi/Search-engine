<script>
    export let label;
  </script>
  
  <style>
    .button {
      width: 120px;
      height: 40px;
      border-radius: 20px;
      border: 2px solid yellow;
      background-color: #fce883;
      font: black;
      font-weight: bold;
      cursor: pointer;
      transition: background-color 0.3s, color 0.3s;
    }
  
    .button:hover {
        background-color:#fedf00;
      font : white;
    }
  </style>
  
  <button class="button">{label}</button>
  