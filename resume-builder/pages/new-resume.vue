<template>
  <div  class="p-12">
    <div class="font-bold mb-10 text-2xl">ساخت رزومه جدید</div>
   <v-expansion-panels>
        <v-expansion-panel class="pa-5">
            <v-expansion-panel-title>اطلاعات کلی</v-expansion-panel-title>
            <v-expansion-panel-text class="mt-5">
                <v-row>
                    <v-col cols="12" sm="4">
                        <v-text-field label="عنوان رزومه" :v-model="resume.resumeTitle" />
                    </v-col>
                </v-row>
                <v-row>
                    <v-col cols="12" sm="4">
                        <v-text-field label="عنوان شغلی" :v-model="resume.jobTitle" />
                    </v-col>
                </v-row>
                <v-row>
                    <v-col cols='12' sm="4">
                        <v-select :items="resumeTemplates" item-title="description" item-value="id" label="قالب مدنظر شما"/>
                    </v-col>
                </v-row>

            </v-expansion-panel-text>
        </v-expansion-panel>
        <v-expansion-panel class="pa-5" :title="FORM_NAMES.PERSONAL_INFO">
            <v-expansion-panel-text class="mt-4" >
                <v-row>
                <v-col cols="12" sm='4'>
                    <v-text-field :label="FORM_NAMES.FIRST_NAME" :v-model="resume.firstName"/>
                </v-col>
                <v-col cols="12" sm='4'>
                    <v-text-field :label="FORM_NAMES.LAST_NAME" :v-model="resume.lastName"/>
                </v-col>
                <v-col cols="12" sm='4'>
                    <v-text-field :label="FORM_NAMES.AGE" :v-model="resume.age"/>
                </v-col>
            </v-row>
            <v-row>
                <v-col cols="12" sm="4">
                    <v-text-field :label="FORM_NAMES.PHONE"  :v-model="resume.phoneNumber" counter maxlength="11" />
                </v-col>
                <v-col cols="12" sm="4">
                    <v-text-field :label="FORM_NAMES.EMAIL"  :v-model="resume.email"/>
                </v-col>
                <v-col cols="12" sm="4">
                    <v-file-input label="عکس پروفایل"  :v-model="resume.image" />
                </v-col>
            </v-row>
            <v-row>
                <v-col cols="12">
                    <v-textarea  :label="FORM_NAMES.ABOUT_ME" :v-model="resume.aboutMe" clearable :hint="FORM_NAMES.ABOUT_ME_PLACEHOLDER" counter />
                </v-col>
            </v-row>
            </v-expansion-panel-text>
        </v-expansion-panel>
        <v-expansion-panel class="pa-5" :title="FORM_NAMES.EDUCATIONAL_INFO">
            <v-expansion-panel-text class="mt-4">
                <v-row v-for="(education,index) in resume.educations" :key="index" class="border-b">
                    <v-col cols=12 sm="3"> 
                        <v-select :items="educationGrades" label="مقطع تحصیلی" :v-model="education.grade"/>
                    </v-col>
                    <v-col cols=12 sm="3"> 
                        <v-text-field  label="نام دانشگاه" :v-model="education.university"/>
                    </v-col>
                    <v-col cols=12 sm="3"> 
                        <v-text-field  label="سال آغار" :v-model="education.start" />
                    </v-col>
                    <v-col cols=12 sm="3"> 
                        <v-text-field label="سال پایان" :v-model="education.end"/>
                    </v-col>
                </v-row>
                <v-row class="mt-5">
                    <v-col cols=12 sm=3>
                        <v-btn variant="outlined" @click="addNewEducationalHistory">
                            اضافه سابقه جدید
                        </v-btn>
                    </v-col>
                </v-row>
            </v-expansion-panel-text>
        </v-expansion-panel>
        <v-expansion-panel title="سوابق کاری" class="pa-5">
            <v-expansion-panel-text>
                <div v-for="(workExperience,index) in resume.workExperiences" :key="index" class="border-b">
                    <v-row class="mt-5">
                        <v-col cols=12 sm="3">
                            <v-text-field label="عنوان شغلی" :v-model="workExperience.title" />
                        </v-col>
                    </v-row>
                    <v-row>
                        <v-col cols="12" sm="4">
                            <v-text-field label="نام شرکت" :v-model="workExperience.company" />
                        </v-col>
                        <v-col cols="6" sm="2">
                            <v-text-field label="سال شروع" :v-model="workExperience.start" />
                        </v-col>
                        <v-col cols="6" sm="2">
                            <v-text-field label="سال پایان" :v-model="workExperience.end" />
                        </v-col>
                    </v-row>
                    <v-row>
                        <v-col cols="12">
                            <v-textarea label="درباره سابقه" :v-model="workExperience.description" clearable counter />
                        </v-col>
                    </v-row>
                </div>
                <v-row class="mt-5">
                    <v-col cols=12 sm=3>
                        <v-btn variant="outlined" @click="addNewWorkExperience">
                            اضافه سابقه جدید
                        </v-btn>
                    </v-col>
                    </v-row>
            </v-expansion-panel-text>
        </v-expansion-panel>
        <v-expansion-panel class="pa-5">
            <v-expansion-panel-title>تکنولوژی ها</v-expansion-panel-title>
            <v-expansion-panel-text class="mt-5">
                <v-row v-for="(skill,index) in resume.skills" :key="index" class="border-b mt-5">
                    <v-col cols=12 sm="3">
                        <v-text-field label="توانایی" :v-model="skill.title"  />
                    </v-col>
                    <v-col cols=12 sm="3">
                        <v-text-field label="میزان تسلط" :v-model="skill.amount" hint="در بازه ۱ تا ۱۰۰"  />
                    </v-col>
                </v-row>
                <v-row>
                    <v-col cols=12 sm=3>
                        <v-btn variant="outlined" @click="addNewSkill">
                            اضافه توانایی جدید
                        </v-btn>
                    </v-col>
                </v-row>
            </v-expansion-panel-text>
        </v-expansion-panel>
        <v-expansion-panel class="pa-5">
            <v-expansion-panel-title>راه های ارتباطی</v-expansion-panel-title>
            <v-expansion-panel-text class="mt-5">
                <v-row v-for="media,index in resume.socialMedias" :key="index" class="border-b">
                    <v-col cols="6" sm="3">
                        <v-select :items="socialMedias" label="راه ارتباطی" :v-model="media.media" />
                    </v-col>
                    <v-col cols="6" sm="3">
                        <v-text-field label="آدرس" :v-model="media.url" />
                    </v-col>
                </v-row>
                <v-row>
                    <v-col cols=12 sm=3>
                        <v-btn variant="outlined" @click="addNewSocialMedia">
                            اضافه راه ارتباطی جدید
                        </v-btn>
                    </v-col>
                </v-row>
            </v-expansion-panel-text>
        </v-expansion-panel>
   </v-expansion-panels>
  </div>
</template>

<script>
export default {
    data(){
        return {
            FORM_NAMES : {
                PERSONAL_INFO : 'اطلاعات شخصی',
                EDUCATIONAL_INFO: 'اطلاعات تحصیلی',
                EMAIL: 'ایمیل',
                PHONE: 'شماره همراه',
                AGE: 'سن',
                LOCATION: 'محل سکونت',
                FIRST_NAME: 'نام',
                LAST_NAME: 'نام خانوادگی',
                ABOUT_ME: 'درباره من',
                ABOUT_ME_PLACEHOLDER : 'چند خطی درباره من',
            },
            resume :{
                resumeTitle: '',
                resumeTemplate: '',
                jobTitle: '',
                firstName: '',
                lastName: '',
                age: '',
                email: '',
                phoneNumber: '',
                location: '',
                image: '',
                aboutMe: 'fffffff',
                educations:[{
                    grade: '',
                    university: '',
                    start: '',
                    end: ''
                }],
                workExperiences:[{
                    title: '',
                    company: '',
                    start: '',
                    end: '',
                    description: ''
                }],
                skills:[{
                    title: '',
                    amount: ''
                }],
                socialMedias:[{
                    media: '',
                    url: ''
                }]
            },
            validators:{
                required: value => !!value || 'Required.'
            },
            educationGrades: ['کارشناسی','کارشناسی ارشد','دکتری'],
            socialMedias: ['instagram', 'linkedin', 'github'],
            resumeTemplates: [
                    {
                        id:1,
                        description: 'قالب اول'
                    },
                    {
                        id:2,
                        description: 'قالب دوم'
                    },
                    {
                        id:3,
                        description: 'قالب سوم'
                    }
                ]
        }
    },
    methods:{
        addNewEducationalHistory(){
            this.resume.educations.push(
                {
                    grade: '',
                    university: '',
                    start: '',
                    end: ''
                }
            )
        },
        addNewWorkExperience(){
            this.resume.workExperiences.push({
                    title: '',
                    company: '',
                    start: '',
                    end: '',
                    description: ''
                })
        },
        addNewSkill(){
            this.resume.skills.push({
                title:'',
                amount: ''
            })
        },
        addNewSocialMedia(){
            this.resume.socialMedias.push({
                media: '',
                url: ''
            })
        },
        submit(){
            console.log(this.resume)
        }
    }
}
</script>

<style>

</style>