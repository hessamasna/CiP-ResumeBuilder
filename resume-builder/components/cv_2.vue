<template>
  <div class="pa-12">
    <div class="head">
      <div class="head_top text-white p-5 text-center bg-neutral-700">
        <div class="head_top_name border-2 border-white p-5 mb-3 font-bold">
         {{ personal.name }}
        </div>
        <span>{{ personal.cvTitle }}</span>
      </div>
      <div class="head_info bg-orange-200	p-5 text-center flex justify-between">
        <div class="w-1/4">
          <div class="font-bold	border-b-2 border-black">محل زندگی</div>
          <div class="mt-1">{{ personal.address }}</div>
        </div>
        <div class="w-1/4 mr-4">
          <div class="font-bold	border-b-2 border-black">سن</div>
          <div class="mt-1">{{ personal.age }}</div>
        </div>
        <div class="w-1/4 mr-4">
          <div class="font-bold	border-b-2 border-black">شماره تلفن</div>
          <div class="mt-1">{{ personal.phone }}</div>
        </div>
        <div class="w-1/4 mr-4">
          <div class="font-bold	border-b-2 border-black">ایمیل</div>
          <div class="mt-1">{{ personal.email }}</div>
        </div>
      </div>
    </div>
    <div class="details mt-5 flex justify-between">
      <div class="details_left w-2/3 ml-2">
        <div class="details_section">
          <div class="details_section_title bg-orange-200 font-bold text-center">
            درباره من
          </div>
          <div class="details_section_details p-4 text-justify">
            {{ aboutMe }}
          </div>
        </div>
        <div class="details_section">
          <div class="details_section_title bg-orange-200 font-bold text-center">
            سوابق شغلی
          </div>
          <div class="details_section_details p-2">
            <div v-for="(item,index) in workHistory" :key="index" class="mb-2">
              <div class="font-bold">{{ item.title }}</div>
              <div>{{ item.place }}</div>
              <div class="text-slate-600">
                {{ item.startDate }} - {{ item.endDate }}
              </div>
              <div>
                {{ item.description }}
              </div>
            </div>
          </div>
        </div>
        
        
      </div>
      <div class="details_right w-1/3">
        <div class="details_section">
          <div class="details_section_title bg-orange-200 font-bold text-center">
            تکنولوژی ها
          </div>
          <div class="details_section_details p-2">
            <div v-for="(item,index) in skills" :key="index" class="flex items-center justify-between">
              <div>{{ item.title }}</div>
              <v-rating v-model="item.amount" density="compact" />
            </div>
          </div>
        </div>
        <div class="details_section">
          <div class="details_section_title bg-orange-200 font-bold text-center">
            سوابق تحصیلی
          </div>
          <div class="details_section_details p-2">
            <div v-for="(item,index) in educations" :key="index" class="mb-2">
              <div class="font-bold">{{ item.degree }}</div>
              <div class="font-semibold	">{{ item.title }}</div>
              <div>{{ item.place }}</div>
              <div class="text-slate-600">
                {{ item.startDate }} - {{ item.endDate }}
              </div>
              <div>
                {{ item.description }}
              </div>
            </div>
          </div>
        </div>
        <div class="details_section">
          <div class="details_section_title bg-orange-200 font-bold text-center">
            راه های ارتباطی
          </div>
          <div class="details_section_details pa-2 flex">
            <div v-for="(media,index) in personal.social" :key="index" class="ml-2">
              <v-icon :icon="media.icon" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name :'CV2',
data() {
    return {
      image: "https://tailone.tailwindtemplate.net/src/img/dummy/avatar1.png",
      personal: {
        name: "سیپ سی وی",
        cvTitle: "توسعه دهنده",
        phone: "09121112233",
        age:22,
        email: "Cv@cipCv.com",
        website: "www.cipCv.com",
        address: "تهران",
        social: [
          {
            title: "linkedin",
            path: "CipCV",
            url: "https://linkedin.com/cipC",
            icon: "mdi-linkedin"
          },
          {
            title: "instagram",
            path: "CipCV",
            url: "https://instagram.com/CipCv",
            icon: "mdi-instagram"
          },
        ]
      },
      aboutMe: "لورم ایپسوم متن ساختگی با تولید سادگی نامفهوم از صنعت چاپ، و با استفاده از طراحان گرافیک است، چاپگرها و متون بلکه روزنامه و مجله در ستون و سطرآنچنان که لازم است، و برای شرایط فعلی تکنولوژی مورد نیاز، و کاربردهای متنوع با هدف بهبود ابزارهای کاربردی می باشد، کتابهای زیادی در شصت و سه درصد گذشته حال و آینده، شناخت فراوان جامعه و متخصصان را می طلبد، تا با نرم افزارها شناخت بیشتری را برای طراحان رایانه ای علی الخصوص طراحان خلاقی، و فرهنگ پیشرو در زبان فارسی ایجاد کرد، در این صورت می توان امید داشت که تمام و دشواری موجود در ارائه راهکارها، و شرایط سخت تایپ به پایان رسد و زمان مورد نیاز شامل حروفچینی دستاوردهای اصلی، و جوابگوی سوالات پیوسته اهل دنیای موجود طراحی اساسا مورد استفاده قرار گیرد.",
      workHistory: [
        {
          title: 'FrontEnd Developer',
          startDate: '2021',
          endDate: '2022',
          place: "Sharif",
          description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolorem explicabo impedit magni reiciendis suscipit! Adipisci deleniti earum eveniet incidunt ipsa libero modi, molestiae nobis pariatur quod repellat, repudiandae rerum sit!"
        },
        {
          title: 'توسعه دهنده فرانت',
          startDate: '1399',
          endDate: '1400',
          place: "دانشگاه صنعتی شریف",
          description: " زمان مورد نیاز شامل حروفچینی دستاوردهای اصلی، و جوابگوی سوالات پیوسته اهل دنیای موجود طراحی اساسا مورد استفاده قرار گیرد."
        }
      ],
      skills: [
        {
          title: "Java",
          amount: 4,
        },
        {
          title: 'JS',
          amount: 3.5,
        },
        {
          title: 'Ruby',
          amount: 2,
        }
      ],
      educations: [
        {
          title: "مهندسی کامپیوتر",
          place: "صنعتی شریف",
          degree: "کارشناسی",
          startDate: "1398",
          endDate: "1402",
        },
        {
          title: "مهندسی کامپیوتر",
          place: "صنعتی شریف",
          degree: "کارشناسی ارشد",
          startDate: "1402",
          endDate: "1406",
        },
      ]
    }
  }
}
</script>

<style scoped lang="scss">
  .head_top{
    // background-color: black;
  }
</style>