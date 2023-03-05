<template>
  <div >
    <loading class="h-screen" v-if="loading"/>
    <div class="flex flex-row pa-12" :style="'font-family: '+ data.font+ '!important'" v-else>
      <div class="basis-1/3 bg-neutral-800 text-white py-10 pr-5">
        <div class="relative overflow-hidden px-6">
          <img :src="data.image"
               class="max-w-full h-auto mx-auto rounded-full bg-gray-50 grayscale mb-4" alt="title image">
        </div>
        <div>
          <div class="mb-8">
            <span class="title">تحصیلات</span>
            <div class="divider_custom" :style="'background-color: '+data.color"></div>
            <div>
              <div v-for="(education, index) in data.educations" :key="index" class="flex mb-4 mr-6">
                <div>
                  <v-icon>mdi-square-small</v-icon>
                </div>
                <div class="flex flex-column mr-2" :style="'font-size: '+data.fontSize+'px'">

                  <span class="text-lg font-weight-bold" :style="'font-size: '+data.fontSize+'px'">{{ education.title }}</span>
                  <span class="text-sm" :style="'font-size: '+data.fontSize+'px'">{{ education.place }} {{ education.degree }}</span>
                  <span class="text-sm" :style="'font-size: '+data.fontSize+'px'">{{ education.startDate }} - {{ education.endDate }}</span>
                </div>
              </div>
            </div>
          </div>
          <div class="mb-8">
            <span class="title">شبکه های اجتماعی</span>
            <div class="divider_custom" :style="'background-color: '+data.color"></div>
            <div :style="'font-size: '+data.fontSize+'px'">
              <div v-for="(social, index) in data.personal.social" :key="index" class="flex mb-2 mr-6 align-center center">
                <v-icon>mdi-{{ social.title }}</v-icon>
                <NuxtLink :to="social.url" target="_blank" class=" mx-3">
                  {{ social.path }}
                </NuxtLink>
              </div>
            </div>
          </div>
          <div class="align-self-end">
            <div class="mb-2" :style="'font-size: '+data.fontSize+'px'">
              <span class="subtitle" :style="'border-color: '+data.color">شماره تماس:</span>
              {{ data.personal.phone }}
            </div>
            <div class="mb-2" :style="'font-size: '+data.fontSize+'px'">
              <span class="subtitle" :style="'border-color: '+data.color">ایمیل:</span>

              {{ data.personal.email }}
            </div>
            <div class="mb-2" :style="'font-size: '+data.fontSize+'px'">
              <span class="subtitle" :style="'border-color: '+data.color">وبسایت:</span>

              {{ data.personal.website }}
            </div>
            <div class="mb-2" :style="'font-size: '+data.fontSize+'px'">
              <span class="subtitle" :style="'border-color: '+data.color">آدرس:</span>

              {{ data.personal.address }}
            </div>
          </div>
        </div>
      </div>
      <div class="basis-3/4 pb-5">
        <div class="pr-5 py-16" :style="'background-color: '+data.color">
          <div class="text-5xl font-weight-bold" :style="'font-size: '+titleFontSize(data.fontSize) +'px'">{{ data.personal.name }}</div>
          <div class="text-2xl font-weight-thin pt-2 text-stone-700" :style="'font-size: '+data.fontSize+'px'">{{ data.personal.cvTitle }}</div>
        </div>
        <div class="px-5">

          <div>
            <div class="title">درباره من</div>
            <div class="divider_custom" :style="'background-color: '+data.color"></div>
            <div class="px-5" :style="'font-size: '+data.fontSize+'px'">{{ data.aboutMe }}</div>
          </div>
          <div>
            <div class="title">سابقه شغلی</div>
            <div class="divider_custom" :style="'background-color: '+data.color"></div>
            <div v-for="(work,index) in data.workHistory" :key="index" class="px-5 flex mb-4">
              <div>
                <v-icon :color="data.color">mdi-square-small</v-icon>
              </div>
              <div class="flex flex-column mr-2" >
                <div class="text-lg font-weight-bold" :style="'font-size: '+data.fontSize+'px'" >{{ work.title }}</div>
                <div class="text-sm" :style="'font-size: '+data.fontSize+'px'" >{{ work.place }}</div>
                <div class="text-sm" :style="'font-size: '+data.fontSize+'px'">{{ work.startDate }} - {{ work.endDate }}</div>
                <div class="text-sm" :style="'font-size: '+data.fontSize+'px'"> {{ work.description }}</div>
              </div>
            </div>
          </div>
          <div>
            <div class="title">مهارت ها</div>
            <div class="divider_custom" :style="'background-color: '+data.color"></div>
            <div class="grid grid-cols-4 gap-4" >

              <div v-for="(skill,index) in data.skills" :key="index" class="px-5 ">
                <div class="font-weight-bold" :style="'font-size: '+data.fontSize+'px'">
                  {{ skill.title }}
                </div>
                <div>
                  <v-progress-linear :color="data.color" :model-value="skill.amount"></v-progress-linear>
                </div>
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
  methods:{
    titleFontSize(int){
      return int+10;
    },
  },
  data() {
    return {
      // data: {
      //   image: "https://tailone.tailwindtemplate.net/src/img/dummy/avatar1.png",
      //   personal: {
      //     name: "سیپ سی وی",
      //     cvTitle: "توسعه دهنده",
      //     phone: "09121112233",
      //     email: "Cv@cipCv.com",
      //     website: "www.cipCv.com",
      //     address: "Tehran",
      //     social: [
      //       {
      //         title: "linkedin",
      //         path: "CipCV",
      //         url: "https://linkedin.com/cipC"
      //       },
      //       {
      //         title: "instagram",
      //         path: "CipCV",
      //         url: "https://instagram.com/CipCv"
      //       },
      //     ]
      //   },
      //   aboutMe: "لورم ایپسوم متن ساختگی با تولید سادگی نامفهوم از صنعت چاپ، و با استفاده از طراحان گرافیک است، چاپگرها و متون بلکه روزنامه و مجله در ستون و سطرآنچنان که لازم است، و برای شرایط فعلی تکنولوژی مورد نیاز، و کاربردهای متنوع با هدف بهبود ابزارهای کاربردی می باشد، کتابهای زیادی در شصت و سه درصد گذشته حال و آینده، شناخت فراوان جامعه و متخصصان را می طلبد، تا با نرم افزارها شناخت بیشتری را برای طراحان رایانه ای علی الخصوص طراحان خلاقی، و فرهنگ پیشرو در زبان فارسی ایجاد کرد، در این صورت می توان امید داشت که تمام و دشواری موجود در ارائه راهکارها، و شرایط سخت تایپ به پایان رسد و زمان مورد نیاز شامل حروفچینی دستاوردهای اصلی، و جوابگوی سوالات پیوسته اهل دنیای موجود طراحی اساسا مورد استفاده قرار گیرد.",
      //   workHistory: [
      //     {
      //       title: 'FrontEnd Developer',
      //       startDate: '2021',
      //       endDate: '2022',
      //       place: "Sharif",
      //       description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolorem explicabo impedit magni reiciendis suscipit! Adipisci deleniti earum eveniet incidunt ipsa libero modi, molestiae nobis pariatur quod repellat, repudiandae rerum sit!"
      //     },
      //     {
      //       title: 'توسعه دهنده فرانت',
      //       startDate: '1399',
      //       endDate: '1400',
      //       place: "دانشگاه صنعتی شریف",
      //       description: " زمان مورد نیاز شامل حروفچینی دستاوردهای اصلی، و جوابگوی سوالات پیوسته اهل دنیای موجود طراحی اساسا مورد استفاده قرار گیرد."
      //     }
      //   ],
      //   skills: [
      //     {
      //       title: "Java",
      //       amount: 80,
      //     },
      //     {
      //       title: 'JS',
      //       amount: 85,
      //     },
      //     {
      //       title: 'Ruby',
      //       amount: 20,
      //     }
      //   ],
      //   educations: [
      //     {
      //       title: "مهندسی کامپیوتر",
      //       place: "صنعتی شریف",
      //       degree: "کارشناسی",
      //       startDate: "1398",
      //       endDate: "1402",
      //     },
      //     {
      //       title: "مهندسی کامپیوتر",
      //       place: "صنعتی شریف",
      //       degree: "کارشناسی ارشد",
      //       startDate: "1402",
      //       endDate: "1406",
      //     },
      //   ]
      // }
    }
  },
}
</script>
<style lang="scss" scoped>
.title {
  font-size: 1.5rem;
  font-weight: bolder;
  margin-top: 1rem;
}

.divider_custom {
  background-color: orange;
  height: 4px;
  margin: 5px 0 15px 0;
}

.subtitle {
  font-weight: bold;
  border-right: solid orange 40px;
  padding: 0 8px 0 2px;
}
</style>
