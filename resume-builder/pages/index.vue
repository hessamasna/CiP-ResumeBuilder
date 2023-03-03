<template>
  <div>
    <div class="flex flex-row align-center">
      <div class="flex-1 text-center">
        <div class="flex flex-column">
          <span class="text-9xl text-[#FFC727] font-weight-bold">
            خلق
          </span>
          <span class="text-2xl mr-48 mt-10 font-weight-bold ">
           یک رزومه با کمک هم
          </span>
        </div>
      </div>
      <div class="align-self-end">
        <img src="/img/cv_main.png" alt="main-banner" class=" max-h-[90vh]">
      </div>
    </div>
    <Cv_1/>
    <CV2/>
  </div>
</template>

<script setup>

import CV2 from "../components/cv_2.vue";
import Cv_1 from "../components/cv_1";
</script>

<style scoped>
</style>
