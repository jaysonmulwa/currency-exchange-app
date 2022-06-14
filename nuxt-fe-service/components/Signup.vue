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
      <div class="container px-5 py-24 mx-auto flex justify-center">
        <div
          class="lg:w-1/3 md:w-1/2 bg-gray-100 rounded-lg p-8 flex flex-col w-full mt-10 md:mt-0 relative z-10 shadow-md"
        >
          <h1
            class="sm:text-3xl text-2xl font-medium title-font mb-2 text-gray-900 tracking"
          >
            Sign Up
          </h1>
          <div class="relative mb-4">
            <label for="username" class="leading-7 text-sm text-gray-600"
              >Username</label
            >
            <input
              v-model="username"
              type="text"
              id="username"
              name="username"
              class="w-full bg-white rounded border border-gray-300 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out"
            />
          </div>

          <div class="relative mb-4">
            <label for="email" class="leading-7 text-sm text-gray-600"
              >Email</label
            >
            <input
              v-model="email"
              type="email"
              id="email"
              name="email"
              class="w-full bg-white rounded border border-gray-300 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out"
            />
          </div>

          <div class="relative mb-4">
            <label for="password" class="leading-7 text-sm text-gray-600"
              >Choose Password</label
            >
            <input
              v-model="password"
              type="password"
              id="password"
              name="password"
              class="w-full bg-white rounded border border-gray-300 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out"
            />
          </div>

          <div class="relative mb-4">
            <label for="firstname" class="leading-7 text-sm text-gray-600"
              >First name</label
            >
            <input
              v-model="firstname"
              type="text"
              id="firstname"
              name="firstname"
              class="w-full bg-white rounded border border-gray-300 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out"
            />
          </div>

          <div class="relative mb-4">
            <label for="lastname" class="leading-7 text-sm text-gray-600"
              >Last name</label
            >
            <input
              v-model="lastname"
              type="text"
              id="lastname"
              name="lastname"
              class="w-full bg-white rounded border border-gray-300 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out"
            />
          </div>

          <button
            class="text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded text-lg"
            @click="signup"
          >
            Sign up
          </button>
          <p class="text-xs text-gray-500 mt-3">
            Already have an account?.
            <NuxtLink to="/login"
              ><b><u>Log in</u></b></NuxtLink
            >
            here.
          </p>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Signup',
  data() {
    return {
      username: '',
      email: '',
      password: '',
      firstname: '',
      lastname: '',
    }
  },
  methods: {
    signup() {
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
        username: this.username,
        email: this.email,
        password: this.password,
        firstname: this.firstname,
        lastname: this.lastname,
        default_currency: "USD",
      }

      const BASE_URL = 'http://localhost:3001'

      axios
        .create(params)
        .post(`${BASE_URL}/api/v1/signup`, payload)
        .then((response) => {
          this.$toast.show('Successfully signed up!')
          console.log(response.data)
          localStorage.setItem('user_id', response.data.data.user_id)
          setInterval(() => {
            this.$router.push('/home')
          }, 1000)
        })
        .catch((error) => {
          this.$toast.show('Something went wrong')
        })
    },
  },
}
</script>
