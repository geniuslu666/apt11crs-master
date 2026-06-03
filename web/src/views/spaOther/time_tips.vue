<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        title=""
      >
        <n-form
          ref="formRef"
          :model="formValue"
          :rules="rules"
          :label-placement="settingStore.isMobile ? 'top' : 'left'"
          :label-width="160"
          class="py-4"
        >
          <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
            <n-gi span="1">
              <n-form-item label="服务时段提示_简体中文" path="orderTimeTips_zh">
                <n-input type="textarea" placeholder="提示信息" v-model:value="formValue.orderTimeTips_zh" @blur="translate('orderTimeTips','zh',formValue.orderTimeTips_zh)" style="width: 400px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="服务时段提示_English" path="orderTimeTips_en">
                <n-input type="textarea" placeholder="提示信息" v-model:value="formValue.orderTimeTips_en" @blur="translate('orderTimeTips','en',formValue.orderTimeTips_en)" style="width: 400px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="服务时段提示_日本语" path="orderTimeTips_ja">
                <n-input type="textarea" placeholder="提示信息" v-model:value="formValue.orderTimeTips_ja" style="width: 400px"/>
              </n-form-item>
            </n-gi>
          </n-grid>
          <div style="text-align: center">
            <n-space justify="center">
              <n-button type="info" :loading="formBtnLoading" @click="formSubmit">
                确定
              </n-button>
            </n-space>
          </div>
        </n-form>
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

  const rules = ref({});
  const group = ref('spaothersetting');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const formBtnLoading = ref(false);
  const settingStore = useProjectSettingStore();

  const formValue = ref({
    orderTimeTips_zh: '',
    orderTimeTips_en: '',
    orderTimeTips_ja: '',
    orderTimeTips_ko: '',
    orderTimeTips_zh_CN: '',
  });

  function translate(type,lang,text){
    if(type == 'orderTimeTips'){
      if(lang == 'zh'){
        // 简体中文
        formValue.value.orderTimeTips_zh_CN = ''
      }else if(lang == 'en'){
        // 韩语
        formValue.value.orderTimeTips_ko = ''
      }
    }
    if(text){
      if(lang == 'en'){
        text = text.toLowerCase()
      }
      Text({
        text: text
      }).then((_res) => {
        if(type == 'orderTimeTips'){
          if(lang == 'zh'){
            // 简体中文
            formValue.value.orderTimeTips_zh_CN = _res.zh_CN
          }else if(lang == 'en'){
            // 韩语
            formValue.value.orderTimeTips_ko = _res.ko
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
            formValue.value.orderTimeTips_zh = res.list.orderTimeTips_zh;
            formValue.value.orderTimeTips_en = res.list.orderTimeTips_en;
            formValue.value.orderTimeTips_ja = res.list.orderTimeTips_ja;
            formValue.value.orderTimeTips_ko = res.list.orderTimeTips_ko;
            formValue.value.orderTimeTips_zh_CN = res.list.orderTimeTips_zh_CN;
          }
          show.value = false;
        }).catch((err)=>{
          show.value = false;
        })
    });
  }
</script>
