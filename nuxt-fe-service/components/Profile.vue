<!-- Please remove this file from your project -->
<template>
  <section class="text-gray-600 body-font">
    <div class="container px-5 pt-12 mx-auto">
      <div class="flex flex-wrap w-full flex-row justify-between">
        <span>
          <NuxtLink to="/home">
            <div
              class="w-10 h-10 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-4"
            >
              <svg
                fill="none"
                stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                class="w-6 h-6"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M10 19l-7-7m0 0l7-7m-7 7h18"
                />
              </svg>
            </div>
          </NuxtLink>
        </span>
      </div>
    </div>
    <div class="container px-5 pt-12 pb-24 mx-auto">
      <div
        class="flex flex-wrap w-full mb-20 flex-col items-center text-center"
      >
        <h1
          class="sm:text-3xl text-2xl font-medium title-font mb-2 text-gray-900 tracking"
        >
          Profile.
        </h1>
      </div>
      <div class="flex flex-wrap justify-center -m-4">
        <div class="w-full lg:w-3/5 p-4">
          <div
            class="border border-gray-200 p-6 rounded-lg flex flex-wrap justify-center -m-4"
          >
            <span
              class="w-full xl:w-1/2 lg:w-1/2 md:w-1/2 sm:w-full xs:w-full p-4"
            >
              <div class="flex items-center justify-center h-full">
                <div
                  class="w-48 h-48 rounded-full inline-flex items-center justify-center bg-gray-200 text-gray-400"
                >
                  <svg
                    fill="none"
                    stroke="currentColor"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    class="w-24 h-24"
                    viewBox="0 0 24 24"
                  >
                    <path d="M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2"></path>
                    <circle cx="12" cy="7" r="4"></circle>
                  </svg>
                </div>
              </div>
            </span>
            <span
              class="w-full xl:w-1/2 lg:w-1/2 md:w-1/2 sm:w-full xs:w-full p-4"
            >
              <div class="flex border-b border-gray-200 py-2">
                <span class="text-gray-500">Username</span>
                <span class="ml-auto text-gray-900">jj19990001</span>
              </div>

              <div class="flex border-b border-gray-200 py-2">
                <span class="text-gray-500">Default Currency</span>
                <span
                  class="ml-auto text-gray-900 flex flex-row"
                  v-if="!editable"
                >
                  <p class="px-2">USD</p>
                  <small @click="editable = true"
                    ><a><u>Edit</u></a></small
                  >
                </span>
                <span
                  class="ml-auto text-gray-900 flex flex-row"
                  v-else-if="editable"
                >
                  <select
                    @change="selected($event)"
                    v-model="default_currency"
                    id="default_currency"
                    class="block w-full px-4 py-2 bg-white border border-gray-400 rounded-lg shadow-sm focus:outline-none focus:shadow-outline-blue focus:border-blue-300 transition duration-150 ease-in-out sm:text-sm sm:leading-5"
                  >
                    <option value="">-- Select a Currency --</option>
                    <option value="USD">USD</option>
                    <option value="KES">KES</option>
                  </select>
                  <small @click="save" class="pl-2"
                    ><a><u>Save</u></a></small
                  >
                </span>
              </div>

              <div class="flex border-b border-gray-200 py-2">
                <span class="text-gray-500">Email address</span>
                <span class="ml-auto text-gray-900"
                  >jayson.mulwa@gmail.com</span
                >
              </div>

              <div class="flex border-b border-gray-200 py-2">
                <span class="text-gray-500">First name</span>
                <span class="ml-auto text-gray-900">Jayson</span>
              </div>

              <div class="flex border-b border-gray-200 py-2">
                <span class="text-gray-500">Last name</span>
                <span class="ml-auto text-gray-900">Mulwa</span>
              </div>
            </span>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Profile',
  data() {
    return {
      username: '',
      email: '',
      first_name: '',
      last_name: '',
      default_currency: '',
      editable: false,
      selected_currency: '',
      profile: {},
    }
  },
  methods: {
    updateInfo() {
      axios
        .get('/api/user')
        .then((response) => {
          this.user = response.data
        })
        .catch((error) => {
          console.log(error)
        })
    },
    save() {

      if (this.default_currency == '') {
        this.$toast.show("Please select a currency")
        return
      }

      const params = {
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
          'Access-Control-Allow-Origin': '*',
          'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, OPTIONS',
          'Access-Control-Allow-Headers': '*'
        },
      }

      const payload = {
        default_currency: this.selected_currency,
      }

      const BASE_URL = 'http://localhost:3001'
      const userId = localStorage.getItem('userId')

      axios
        .create(params)
        .put(`${BASE_URL}/api/v1/profile/${userId}`, payload)
        .then((response) => {
          console.log(response)
          this.default_currency = this.selected_currency
          this.editable = false
        })
        .catch((error) => {
          console.log(error)
          this.editable = false
        })
    },
    selected(e) {
      console.log(e.target.value)
      this.selected_currency = e.target.value
    },
  },

  async mounted() {
    try {
      const params = {
        withCredentials: true,
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
          'Access-Control-Allow-Origin': '*',
          'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, OPTIONS',
          'Access-Control-Allow-Headers': '*'
        },
      }

      const BASE_URL = 'http://localhost:3001'

      const userId = localStorage.getItem('user_id')

      const profile = await Promise.all([
        axios.create(params).get(`${BASE_URL}/api/v1/profile/${userId}`),
      ])
      
      this.profile = profile?.data;
      
    } catch (error) {
      console.log(error)
      this.$toast.show('Error fetching data')
    }
  },

}
</script>
