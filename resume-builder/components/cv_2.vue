<template>
  <div>
    <loading class="h-screen" v-if="loading"/>
    <div class="pa-12 pt-2" :style="'font-family: '+ data.font_family+ '!important'" v-else>
      <div class="head">
        <div class="head_top text-white p-5 text-center bg-neutral-700">
          <div class="head_top_name border-2 border-white p-5 mb-3 font-weight-bold"
               :style="'font-size: '+titleFontSize(data.font_size)+'px'">
            {{ data.personal_info.first_name }} {{ data.personal_info.last_name }}
          </div>
          <span :style="'font-size: '+data.font_size+'px'">{{ data.job_title }}</span>
        </div>
        <div class="head_info 	p-5 text-center flex justify-between" :style="'background-color: '+data.color">
          <div class="w-1/4">
            <div class="font-bold	border-b-2 border-black">محل زندگی</div>
            <div class="mt-1" :style="'font-size: '+data.font_size+'px'">{{ data.personal_info.address }}</div>
          </div>
          <div class="w-1/4 mr-4">
            <div class="font-bold	border-b-2 border-black">سن</div>
            <div class="mt-1" :style="'font-size: '+data.font_size+'px'">
              {{ !!data.personal_info.age ? data.personal_info.age : 22 }}
            </div>
          </div>
          <div class="w-1/4 mr-4">
            <div class="font-bold	border-b-2 border-black">شماره تلفن</div>
            <div class="mt-1" :style="'font-size: '+data.font_size+'px'">{{ data.personal_info.phone_number }}</div>
          </div>
          <div class="w-1/4 mr-4">
            <div class="font-bold	border-b-2 border-black">ایمیل</div>
            <div class="mt-1" :style="'font-size: '+data.font_size+'px'">{{ data.personal_info.email }}</div>
          </div>
        </div>
      </div>
      <div class="details mt-5 flex justify-between">
        <div class="details_left w-2/3 ml-2">
          <div class="details_section">
            <div class="details_section_title  font-bold text-center" :style="'background-color: '+data.color">
              درباره من
            </div>
            <div class="details_section_details p-4 text-justify" :style="'font-size: '+data.font_size+'px'">
              {{ data.about_me }}
            </div>
          </div>
          <div class="details_section">
            <div class="details_section_title font-bold text-center" :style="'background-color: '+data.color">
              سوابق شغلی
            </div>
            <div class="details_section_details p-2" :style="'font-size: '+data.font_size+'px'">
              <div v-for="(item,index) in data.experience" :key="index" class="mb-2">
                <div class="flex">
                  <v-icon :color="data.color">mdi-square-small</v-icon>
                  <div>
                    <div class="font-bold">{{ item.title }}</div>
                    <div>{{ item.company }}</div>
                    <div class="text-slate-600">
                      {{ item.start }} - {{ item.end }}
                    </div>
                    <div>
                      {{ item.description }}
                    </div>
                  </div>

                </div>


              </div>
            </div>
          </div>


        </div>
        <div class="details_right w-1/3">
          <div class="details_section">
            <div class="details_section_title font-bold text-center" :style="'background-color: '+data.color">
              تکنولوژی ها
            </div>
            <div class="details_section_details p-2" :style="'font-size: '+data.font_size+'px'">
              <div v-for="(item,index) in data.skills" :key="index" class="flex items-center justify-between">
                <div>{{ item.name }}</div>
                <v-rating disabled :color="data.color" :model-value="convertToFive(item.percent)" density="compact"/>
              </div>
            </div>
          </div>
          <div class="details_section">
            <div class="details_section_title font-bold text-center" :style="'background-color: '+data.color">
              سوابق تحصیلی
            </div>
            <div class="details_section_details p-2" :style="'font-size: '+data.font_size+'px'">
              <div v-for="(item,index) in data.education" :key="index" class="mb-2">
                <div class="flex">
                  <v-icon :color="data.color">mdi-square-small</v-icon>
                  <div>
                    <div class="font-bold">{{ item.degree }}</div>
                    <div class="font-semibold	">{{ item.school }}-{{ item.major }}</div>
                    <div>{{ item.place }}</div>
                    <div class="text-slate-600">
                      {{ item.start }} - {{ item.end }}
                    </div>
                  </div>
                </div>

                <!--                <div>-->
                <!--                  {{ item.description }}-->
                <!--                </div>-->
              </div>
            </div>
          </div>
          <div class="details_section">
            <div class="details_section_title  font-bold text-center" :style="'background-color: '+data.color">
              راه های ارتباطی
            </div>
            <div class="details_section_details pa-2 flex">
              <div v-for="(media,index) in data.social_medias" :key="index" class="ml-2">
                <NuxtLink :to="media.link" target="_blank" class="">
                  <v-icon>mdi-{{ media.plat_form }}</v-icon>
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
  name: 'CV2',
  props: ['data', 'loading'],
  methods: {
    convertToFive(num) {
      return num / 20;
    },
    titleFontSize(int) {
      return Number(int) + 10;
    },
  },
  data() {
    return {
      // data:{
      //   image: "https://tailone.tailwindtemplate.net/src/img/dummy/avatar1.png",
      //   personal: {
      //     name: "سیپ سی وی",
      //     cvTitle: "توسعه دهنده",
      //     phone: "09121112233",
      //     age:22,
      //     email: "Cv@cipCv.com",
      //     website: "www.cipCv.com",
      //     address: "تهران",
      //     social: [
      //       {
      //         title: "linkedin",
      //         path: "CipCV",
      //         url: "https://linkedin.com/cipC",
      //         icon: "mdi-linkedin"
      //       },
      //       {
      //         title: "instagram",
      //         path: "CipCV",
      //         url: "https://instagram.com/CipCv",
      //         icon: "mdi-instagram"
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
      //       amount: 4,
      //     },
      //     {
      //       title: 'JS',
      //       amount: 3.5,
      //     },
      //     {
      //       title: 'Ruby',
      //       amount: 2,
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
  }
}
</script>

<style scoped lang="scss">
.head_top {
  // background-color: black;
}
</style>
