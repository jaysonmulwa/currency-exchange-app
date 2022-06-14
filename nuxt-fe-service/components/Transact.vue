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
          Transact.
        </h1>
      </div>
      <div class="container px-5 py-24 mx-auto flex justify-center">
        <div
          class="lg:w-1/3 md:w-1/2 bg-gray-100 rounded-lg p-8 flex flex-col w-full mt-10 md:mt-0 relative z-10 shadow-md"
        >
          <!--h2 class="text-gray-900 text-lg mb-1 font-medium title-font">Feedback</h2>-->
          <p class="leading-relaxed mb-5 text-gray-600">
            Debit and credit your account
          </p>
          <div class="relative mb-4">
            <label for="type" class="leading-7 text-sm text-gray-600"
              >Transaction type</label
            >
            <select
              v-model="type"
              id="type"
              class="block w-full px-4 py-2 bg-white border border-gray-400 rounded-lg shadow-sm focus:outline-none focus:shadow-outline-blue focus:border-blue-300 transition duration-150 ease-in-out sm:text-sm sm:leading-5"
            >
              <option value="">-- Select a type --</option>
              <option value="dr">Debit</option>
              <option value="cr">Credit</option>
            </select>
          </div>
          <div class="relative mb-4">
            <label for="amount" class="leading-7 text-sm text-gray-600"
              >Amount</label
            >
            <input
              v-model="amount"
              type="number"
              id="amount"
              name="amount"
              class="w-full bg-white rounded border border-gray-300 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out"
            />
          </div>
          <button
            class="text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded text-lg"
            @click="transact"
          >
            Transact
          </button>
          <p class="text-xs text-gray-500 mt-3">
            All debits and credits are in your default currency.
          </p>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Transact',
  data() {
    return {
      type: '',
      amount: '',
    }
  },
  methods: {
    transact() {
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
        user_id: localStorage.getItem('user_id'),
        entry: this.type,
        amount: this.amount,
      }

      const BASE_URL = 'http://localhost:3001'

      axios
        .create(params)
        .post(`${BASE_URL}/api/v1/transact`, payload)
        .then((response) => {
          console.log(response)
          if (response.data.success) {
            this.type == 'debit'
              ? this.$toast.show('You have successfully debited your account.')
              : this.$toast.show('You have successfully credited your account.')
          }
        })
        .catch((error) => {
          console.log(error)
          this.$toast.show('Something went wrong!')
        })
    },
  },
}
</script>
