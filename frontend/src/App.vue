<template>
  <!---------------------------------------- Start Nav ----------------------------------------------------------------->
  <nav class="bg-slate-200 w-full flex relative justify-start items-center mx-auto px-8 h-16">
      <!-- logo -->
      <div class="inline-flex">
          <a class="_o6689fn" href="/"
              ><div class="hidden md:block">
                  <svg version="1.0" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="32" height="32"
                  viewBox="0 0 100 100" enable-background="new 0 0 100 100" xml:space="preserve">
                  <path fill="#263238" d="M50,10H10v40c0,22.09,17.91,40,40,40h40V50C90,27.91,72.09,10,50,10z M83.333,83.333H50
                    c-18.379,0-33.333-14.954-33.333-33.333V16.667H50c18.379,0,33.333,14.955,33.333,33.333V83.333z"/>
                  <path fill="#FFB74D" d="M50,23.333c-14.727,0-26.667,11.94-26.667,26.667c0,14.727,11.94,26.667,26.667,26.667
                    S76.667,64.727,76.667,50C76.667,35.273,64.727,23.333,50,23.333z M50,70c-11.048,0-20-8.955-20-20c0-11.044,8.952-20,20-20
                    c11.045,0,20,8.956,20,20C70,61.045,61.045,70,50,70z"/>
                  <path fill="#FFB74D" d="M50,36.667c-7.363,0-13.333,5.97-13.333,13.333c0,7.363,5.97,13.333,13.333,13.333S63.333,57.363,63.333,50
                    C63.333,42.637,57.363,36.667,50,36.667z M50,56.667c-3.682,0-6.667-2.985-6.667-6.667c0-3.682,2.985-6.667,6.667-6.667
                    c3.682,0,6.667,2.985,6.667,6.667C56.667,53.682,53.682,56.667,50,56.667z"/>
                  </svg>
              </div>
          </a>
      </div>
      <div class="hidden sm:block flex-shrink flex-grow-0 justify-start px-2">
        <p class ="font-poppins">Who did it?</p>
      </div>
  </nav>

    <!---------------------------------------- End Nav ----------------------------------------------------------------->
    <!---------------------------------------- Start searchbar ---------------------------------------------------------->

  <form class='max-w-md mx-auto mt-4 border-solid border rounded-lg border-slate-200' id="search-form" @submit.prevent="searchQuery">
    <div class="relative flex items-center w-full h-12 rounded-lg focus-within:shadow-lg bg-white overflow-hidden">
      <button type="submit" form="search-form" class="focus:outline-none">
        <div class="grid place-items-center h-full w-12 text-gray-300">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>  
        </div>
      </button>
        <input
        class="peer h-full w-full outline-none text-sm text-gray-700 pr-2"
        type="text"
        id="search"
        placeholder="Search something.." v-model="searchTerm" /> 
    </div>
  </form>

  <!---------------------------------------- End searchbar ---------------------------------------------------------->
  <!---------------------------------------- Start grid ------------------------------------------------------------->

<div class="grid grid-cols-2 gap-3 pt-3">
  <div class="flex flex-col">
    <div class="sm:mx-0.5 lg:mx-0.5">
      <div class="py-2 inline-block min-w-full sm:px-6 lg:px-8">
        <div class="overflow-y-auto overflow-x-hidden max-h-[22rem]">
          <table class="w-128 fixed-table">
            <thead class="bg-white border-b">
              <tr>
                <th scope="col" class="text-sm font-medium text-gray-900 px-6 py-4 text-center">
                  Subject
                </th>
                <th scope="col" class="text-sm font-medium text-gray-900 px-6 py-4 text-center">
                  From
                </th>
                <th scope="col" class="text-sm font-medium text-gray-900 px-6 py-4 text-center">
                  To
                </th>
              </tr>
            </thead>
            <tbody v-for="(email, index) in emails" :key="email">
              <template v-if="index%2===0">
                <tr v-bind:data-content="email.Content" v-bind="dataAttributes(email.Content, email.Date)" class="bg-white border-b hover:scale-105 hover:cursor-pointer" @click="showContent">
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900  w-28"><div class="w-28 overflow-hidden">{{email.Subject}}</div></td>
                  <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap  w-28">
                    <div class="w-28 overflow-hidden">{{email.From}}</div>
                  </td>
                  <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap overflow:hidden w-28">
                    <div class="w-28 overflow-hidden">{{ email.To }}</div>
                  </td>
                </tr>
              </template>
              <template v-else>
                <tr v-bind:data-content="email.Content" class="bg-gray-100 border-b hover:scale-105 hover:cursor-pointer" @click="showContent">
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900  w-28"><div class="w-28 overflow-hidden" @click.stop>{{email.Subject}}</div></td>
                  <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap  w-28">
                    <div class="w-28 overflow-hidden">{{email.From}}</div>
                  </td>
                  <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap overflow:hidden w-28">
                    <div class="w-28 overflow-hidden">{{ email.To }}</div>
                  </td>
                </tr>
              </template>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>

  <div class="flex flex-col">
    <div class="py-2 inline-block min-w-full sm:px-6 lg:px-8  overflow-y-auto overflow-x-hidden max-h-[22rem]">
      <p>{{ date }}</p>
      <p class="text-left">{{content}}</p>
    </div>
  </div>
</div>
<!---------------------------------------- End grid ------------------------------------------------------------->
</template>
<script>
  import axios from 'axios';

  export default {
    name: 'App',

    data() { return {
      searchTerm: '',
      emails: [],
      content: "",
      date: ""
    } },

    methods: {
      searchQuery() {
        axios.get(window.location.origin + "/search", {
          params: {
            searchTerm: this.searchTerm
          }
        })
        .then((response) => {
          console.log(response.data)
          this.emails = response.data
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        })
      },
      showContent(event){
        this.content = event.currentTarget.dataset.content
        this.date = event.currentTarget.dataset.date
      },
      dataAttributes(content, date){
        return {
          "data-content": content,
          "data-date": date
        }
      }
    }
  }
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>