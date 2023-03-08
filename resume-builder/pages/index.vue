<template>
  <div class="">
    <div class="flex flex-row align-center hidden md:visible">
      1
      <div class="flex-1 text-center">
        <div class="flex flex-column mt-8">
          <span class="text-5xl text-[#FFC727] font-weight-bold md:text-9xl">
            خلق
          </span>
          <span class="text-2xl md:mr-48 md:mt-10 mt-4 font-weight-bold ">
           یک رزومه با کمک هم
          </span>
        </div>
      </div>
      <div class="align-self-end">
        <img src="/img/cv_main.png" alt="main-banner" class=" max-h-[90vh]">
      </div>
    </div>
    <div class="flex flex-column align-center md:hidden">
      2
      <div class="flex-1 text-center">
        <div class="flex flex-column mt-8">
          <span class="text-5xl text-[#FFC727] font-weight-bold md:text-9xl">
            خلق
          </span>
          <span class="text-2xl md:mr-48 md:mt-10 mt-4 font-weight-bold ">
           یک رزومه با کمک هم
          </span>
        </div>
      </div>
      <div class="align-self-end">
        <img src="/img/cv_main.png" alt="main-banner" class=" max-h-[90vh]">
      </div>
    </div>
<!--    <Cv_1/>-->
<!--    <CV2/>-->
  </div>
</template>

<script setup>

import CV2 from "../components/cv_2.vue";
import Cv_1 from "../components/cv_1";
</script>

<style scoped>
</style>
