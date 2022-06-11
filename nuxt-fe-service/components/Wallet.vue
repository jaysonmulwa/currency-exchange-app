<!-- Please remove this file from your project -->
<template>
  <section class="text-gray-600 body-font">
    <div class="container px-5 pt-12 mx-auto">
      <div class="flex flex-wrap w-full flex-row justify-between">
        <span>
          <NuxtLink to="/home">
          <div class="w-10 h-10 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-4">
            <svg fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" class="w-6 h-6" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
            </svg>
          </div>
          </NuxtLink>
        </span>

      </div>
    </div>
    <div class="container px-5 pt-12 pb-24 mx-auto">
    <div class="flex flex-wrap w-full mb-20 flex-col items-center text-center">
      <h1 class="sm:text-3xl text-2xl font-medium title-font mb-2 text-gray-900 tracking">Wallet.</h1>
    </div>
    <div class="flex flex-wrap justify-center -m-4">
      <div class="xl:w-1/2 md:w-1/2 w-full p-4">
          <div class="border border-gray-200 p-6 rounded-lg">
            <p class="text-lg w-full mt-2">BALANCE</p>
            <p class="text-5xl text-gray-900 font-bold title-font mt-6">$ 4,000.00</p>
            <p class="leading-relaxed text-base w-full mt-2">Currency: United States dollar</p>
          </div>
      </div>
    </div>
    <div class="flex flex-wrap justify-center">
      <div class="container px-5 py-24 mx-auto">
        <div class="lg:w-4/5 mx-auto flex flex-wrap">
          <div class="lg:w-full w-full lg:pr-10 lg:py-6 mb-6 lg:mb-0">
            <div class="flex mb-4">
              <a class="flex-grow text-indigo-500 border-b-2 border-indigo-500 py-2 text-lg px-1">Recent Transaction history</a>
            </div>
            <p class="leading-relaxed mb-4">These are your most recent Credit, Debit and Cash Transfer transactions. For the full transaction statement, kindly get in touch with us.</p>
            <div class="flex border-t border-gray-200 py-2">
              <span class="text-gray-500">Credit </span>
              <span class="ml-auto text-gray-900">$ 2000</span>
            </div>
            <div class="flex border-t border-gray-200 py-2">
              <span class="text-gray-500">Credit</span>
              <span class="ml-auto text-gray-900">$ 2000</span>
            </div>
            <div class="flex border-t border-gray-200 py-2">
              <span class="text-gray-500">Transfer</span>
              <span class="ml-auto text-gray-900">$ 2000</span>
            </div>
            <div class="flex border-t border-gray-200 py-2">
              <span class="text-gray-500">Credit</span>
              <span class="ml-auto text-gray-900">$ 2000</span>
            </div>
            <div class="flex border-t border-gray-200 py-2">
              <span class="text-gray-500">Debit</span>
              <span class="ml-auto text-gray-900">$ 2000</span>
            </div>
            <div class="flex border-t border-b mb-6 border-gray-200 py-2">
              <span class="text-gray-500">Debiit</span>
              <span class="ml-auto text-gray-900">$ 2000</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  </section>
</template>

<script>
export default {
  name: 'Wallet',
}
</script>
