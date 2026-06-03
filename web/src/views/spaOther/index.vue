<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">其他配置</text>
        </template>
        <n-tabs type="card" animated>
          <n-tab-pane name="下单须知" tab="下单须知">
            <BookingNotice ref="bookingNoticeRef"/>
          </n-tab-pane>
          <n-tab-pane name="服务时段提示" tab="服务时段提示">
            <TimeTips ref="timeTipsRef"/>
          </n-tab-pane>

        </n-tabs>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import {useProjectSettingStore} from "@/store/modules/projectSetting";
  import {Text} from "@/api/translate";
  import BookingNotice from "@/views/spaOther/booking_notice.vue";
  import TimeTips from "@/views/spaOther/time_tips.vue";

  const rules = ref({});
  const group = ref('spaothersetting');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const formBtnLoading = ref(false);
  const settingStore = useProjectSettingStore();
  const bookingNoticeRef = ref();
  const timeTipsRef = ref();

  const formValue = ref({
    bookingNotice_zh: '',
    bookingNotice_en: '',
    bookingNotice_ja: '',
    bookingNotice_ko: '',
    bookingNotice_zh_CN: '',
  });

  function translate(type,lang,text){
    if(type == 'bookingNotice'){
      if(lang == 'zh'){
        // 简体中文
        formValue.value.bookingNotice_zh_CN = ''
      }else if(lang == 'en'){
        // 韩语
        formValue.value.bookingNotice_ko = ''
      }
    }
    if(text){
      if(lang == 'en'){
        text = text.toLowerCase()
      }
      Text({
        text: text
      }).then((_res) => {
        if(type == 'bookingNotice'){
          if(lang == 'zh'){
            // 简体中文
            formValue.value.bookingNotice_zh_CN = _res.zh_CN
          }else if(lang == 'en'){
            // 韩语
            formValue.value.bookingNotice_ko = _res.ko
          }
        }
      });
    }
  }

  function formSubmit() {
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {

        updateConfig({ group: group.value, list: formValue.value }).then((_res) => {
          formBtnLoading.value = false;
          message.success('更新成功');
          load();
        }).catch((err) => {
          formBtnLoading.value = false;
        });
      } else {
        message.error('验证失败，请填写完整信息');
        formBtnLoading.value = false;
      }
    });
  }

  onMounted(() => {
    load();
  });

  function load() {
    show.value = true;
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          if(res.list){
            formValue.value = res.list;
          }
          show.value = false;
        }).catch((err)=>{
          show.value = false;
        })
    });
  }
</script>
