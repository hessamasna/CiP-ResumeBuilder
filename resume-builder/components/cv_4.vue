<template>
  <loading class="h-screen" v-if="loading"/>
  <div
      :style="'font-size: '+data.font_size+'px;font-family: '+ data.font_family+ '!important;max-width:90% ;'"
      class=" mx-auto"
      v-else>
    <div class="flex flex-row pa-12">
      <div class="greenSide basis-1/4 text-white py-10 pr-5" :style="'background-color: '+data.color">
        <div class="relative overflow-hidden px-6 mt-12">
          <v-img lazy-src="https://tailone.tailwindtemplate.net/src/img/dummy/avatar1.png"
                 :src="!!data.image?data.image:'https://tailone.tailwindtemplate.net/src/img/dummy/avatar1.png'"
                 class=" w-56 h-auto mx-auto rounded-full bg-gray-50 grayscale mb-4" alt="title image">
            <template v-slot:placeholder>
              <div class="d-flex align-center justify-center fill-height">
                <v-progress-circular
                    color="green"
                    indeterminate
                ></v-progress-circular>
              </div>
            </template>
          </v-img>
        </div>
        <div>
          <div class="mb-10">
            <span class="title">درباره من</span>
            <div class="px-5" :style="'font-size: '+data.font_size+'px;text-align: start;'">{{ data.about_me }}</div>
          </div>

          <div class="align-self-end">
            <span class="title ">اطلاعات تماس</span>
            <div class="mb-2 mt-4" :style="'font-size: '+data.font_size+'px;'">
                            <span>
                                <v-icon color="white" icon="mdi-phone" size="x-small"></v-icon>
                            </span>
              {{ data.personal_info.phone_number }}
            </div>
            <div class="mb-2" :style="'font-size: '+data.font_size+'px;'">
                            <span>
                                <v-icon color="white" icon="mdi-mail" size="x-small"></v-icon>
                            </span>
              {{ data.personal_info.email }}
            </div>
            <div class="mb-2" :style="'font-size: '+data.font_size+'px;'">
                            <span>
                                <v-icon color="white" icon="mdi-map" size="x-small"></v-icon>
                            </span>
              {{ data.personal_info.address }}
            </div>
          </div>
        </div>
      </div>
      <div class="pb-5 flex-1">

        <div class="pr-5 py-16 pinkBar">
          <div class="name text-5xl font-weight-bold" :style="'font-size: '+titleFontSize(data.font_size)+'px;'">
            {{ data.personal_info.first_name }}
            {{ data.personal_info.last_name }}
          </div>
          <div class="name text-2xl font-weight-thin pt-2 text-stone-700" :style="'font-size: '+data.font_size+'px;'">
            {{ data.job_title }}
          </div>
        </div>


        <div class="px-5 mt-6">
          <div class="mb-6">

            <div class="mb-2" :style="'font-size: '+data.font_size+'px;'">
                            <span>
                                <v-icon :color="data.color" icon="mdi-bookmark" size="small"></v-icon>
                              <!-- real icon name is label_important -->
                            </span>
              <span class="title">تحصیلات</span>
            </div>

            <div :style="'font-size: '+data.font_size+'px;'">
              <div v-for="(education, index) in data.education" class="flex mb-4 mr-6">
                <div class="flex flex-column mr-2">
                  <span class="text-lg font-weight-bold">{{ education.major }}</span>
                  <span class="text-sm">{{ education.school }} {{ education.degree }}</span>
                  <span class="text-sm">{{ education.start }} - {{ education.end }}</span>
                </div>
              </div>
            </div>
          </div>
          <div class="mb-6">
            <div class="mb-2" :style="'font-size: '+data.font_size+'px;'">
                            <span>
                                <v-icon :color="data.color" icon="mdi-bookmark" size="small"></v-icon>
                              <!-- real icon name is label_important -->
                            </span>
              <span class="title">سوابق شغلی</span>
            </div>
            <div v-for="(work, index) in data.experience" class="px-5 flex mb-4">
              <div class="flex flex-column mr-2">

                <div class="text-lg font-weight-bold">{{ work.title }}</div>
                <div class="text-sm">{{ work.company }}</div>
                <div class="text-sm">{{ work.start }} - {{ work.end }}</div>
                <div class="text-sm"> {{ work.description }}</div>
              </div>
            </div>
          </div>
          <div class="mb-6">
            <div class="mb-2">
                            <span>
                                <v-icon :color="data.color" icon="mdi-bookmark" size="small"></v-icon>
                              <!-- real icon name is label_important -->
                            </span>
              <span class="title">مهارت ها</span>
            </div>
            <div class="grid grid-cols-1 gap-4">
              <div class="mr-8 font-weight-thin" v-for="(skill, index) in  data.skills">
                {{ skill.name }}
              </div>
            </div>
          </div>
          <div class="mb-6">
            <div class="mb-2">
                            <span>
                                <v-icon :color="data.color" icon="mdi-bookmark" size="small"></v-icon>
                              <!-- real icon name is label_important -->
                            </span>
              <span class="title">زبان</span>
            </div>
            <div class="grid grid-cols-1 gap-4">
              <div class="mr-8" v-for="(lang) in data.social_medias">
                <NuxtLink :to="lang.link" target="_blank" class=" mx-3">
                  {{ lang.plat_form }}
                </NuxtLink>
              </div>
            </div>
          </div>
        </div>

      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: "cv_1",
  props: ['data', 'loading'],
  methods: {
    titleFontSize(int) {
      return Number(int) + 10;
    },
  },
  data() {
    return {}
  },
}
</script>
<style lang="scss" scoped>
.greenSide {
  margin-top: -4%;
  z-index: 1;
}

.pinkBar {
  background-color: #ebcbbb;
  text-align: center;
}

.title {
  //font-size: 1.5rem;
  font-weight: bolder;
  margin-top: 1rem;
  margin-right: 10px;
}

.name {

}

@media screen and (min-width: 500px) {
  .pinkBar {
    background-color: #ebcbbb;
    text-align: center;
    margin-right: -25%;
  }
  .name {
    margin-right: -40%;
  }
}
</style>
