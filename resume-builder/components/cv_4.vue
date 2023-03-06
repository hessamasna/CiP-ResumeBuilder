<template>
    <div style="max-width:90% ; margin: auto;">
        <div class="flex flex-row pa-12">
            <div class="greenSide basis-1/3 text-white py-10 pr-5">
                <div class="relative overflow-hidden px-6 mt-12">
                    <img :src="image" class="max-w-full h-auto mx-auto rounded-full bg-gray-50 grayscale mb-4 pa-2"
                        alt="title image">
                </div>
                <div>
                    <div class="mb-8">
                        <span class="title">درباره من</span>
                        <div class="px-5" style="text-align: start;">
                            {{ aboutMe }}
                        </div>
                    </div>

                    <div class="align-self-end">
                        <span class="title">اطلاعات تماس</span>
                        <div class="mb-2">
                            <span>
                                <v-icon color="white" icon="mdi-phone" size="x-small"></v-icon>
                            </span>
                            {{ personal.phone }}
                        </div>
                        <div class="mb-2">
                            <span>
                                <v-icon color="white" icon="mdi-mail" size="x-small"></v-icon>
                            </span>
                            {{ personal.email }}
                        </div>
                        <div class="mb-2">
                            <span>
                                <v-icon color="white" icon="mdi-map" size="x-small"></v-icon>
                            </span>
                            {{ personal.address }}
                        </div>
                    </div>
                </div>
            </div>

            <div class="pb-5">

                <div class="pr-5 py-16 pinkBar">
                    <div class="name text-5xl font-weight-bold">{{ personal.name }}</div>
                    <div class="name text-2xl font-weight-thin pt-2 text-stone-700">{{ personal.cvTitle }}</div>
                </div>


                <div class="px-5">
                    <div>

                        <div class="mb-2">
                            <span>
                                <v-icon color="#6c999a" icon="mdi-bookmark" size="small"></v-icon>
                                <!-- real icon name is label_important -->
                            </span>
                            <span class="title">تحصیلات</span>
                        </div>

                        <div>
                            <div v-for="(education, index) in educations" class="flex mb-4 mr-6">
                                <div class="flex flex-column mr-2">
                                    <span class="text-lg font-weight-bold">{{ education.title }}</span>
                                    <span class="text-sm">{{ education.place }} {{ education.degree }}</span>
                                    <span class="text-sm">{{ education.startDate }} - {{ education.endDate }}</span>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div>
                        <div class="mb-2">
                            <span>
                                <v-icon color="#6c999a" icon="mdi-bookmark" size="small"></v-icon>
                                <!-- real icon name is label_important -->
                            </span>
                            <span class="title">سوابق شغلی</span>
                        </div>
                        <div v-for="(work, index) in workHistory" class="px-5 flex mb-4">
                            <div class="flex flex-column mr-2">

                                <div class="text-lg font-weight-bold">{{ work.title }}</div>
                                <div class="text-sm">{{ work.place }}</div>
                                <div class="text-sm">{{ work.startDate }} - {{ work.endDate }}</div>
                                <div class="text-sm"> {{ work.description }}</div>
                            </div>
                        </div>
                    </div>
                    <div>
                        <div class="mb-2">
                            <span>
                                <v-icon color="#6c999a" icon="mdi-bookmark" size="small"></v-icon>
                                <!-- real icon name is label_important -->
                            </span>
                            <span class="title">مهارت ها</span>
                        </div>
                        <div class="grid grid-cols-1 gap-4">
                            <div class="mr-8" v-for="(skill, index) in skills">
                                {{ skill.title }}
                            </div>
                        </div>
                    </div>
                    <div>
                        <div class="mb-2">
                            <span>
                                <v-icon color="#6c999a" icon="mdi-bookmark" size="small"></v-icon>
                                <!-- real icon name is label_important -->
                            </span>
                            <span class="title">زبان</span>
                        </div>
                        <div class="grid grid-cols-1 gap-4">
                            <div class="mr-8" v-for="(lang) in languages">
                                {{ lang.title }}
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
    data() {
        return {
            image: "https://tailone.tailwindtemplate.net/src/img/dummy/avatar1.png",
            personal: {
                name: "سیپ سی وی",
                cvTitle: "توسعه دهنده",
                phone: "09121112233",
                email: "Cv@cipCv.com",
                website: "www.cipCv.com",
                address: "Tehran",
                social: [
                    {
                        title: "linkedin",
                        path: "CipCV",
                        url: "https://linkedin.com/cipC"
                    },
                    {
                        title: "instagram",
                        path: "CipCV",
                        url: "https://instagram.com/CipCv"
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
                    amount: 80,
                },
                {
                    title: 'JS',
                    amount: 85,
                },
                {
                    title: 'Ruby',
                    amount: 20,
                }
            ],
            languages: [
                {
                    title: "انگلیسی"
                },
                {
                    title: "آلمانی"
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
    },
}
</script>
<style lang="scss" scoped>
.greenSide {
    background-color: #6c999a;
    margin-top: -4%;
    z-index: 1;
}

.pinkBar {
    background-color: #ebcbbb;
    text-align: center;
}

.title {
    font-size: 1.5rem;
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
  .name{
    margin-right: -40%;
  }
}
</style>
