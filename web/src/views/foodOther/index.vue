<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">其他配置</text>
        </template>
        <n-form
          ref="formRef"
          :model="formValue"
          :rules="rules"
          :label-placement="settingStore.isMobile ? 'top' : 'left'"
          :label-width="150"
          class="py-4"
          style="padding-top: 0"
        >
          <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
            <n-gi>
              <n-divider title-placement="left">
                排序模式
              </n-divider>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="前端餐厅排序模式" path="sortType">
                <a-radio-group v-model:value="formValue.sortType" name="radioGroup">
                  <a-radio :value="1">按最早可预约时间</a-radio>
                  <a-radio :value="2">按后端排序倒序</a-radio>
                </a-radio-group>
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-divider title-placement="left">
                预定须知
              </n-divider>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="预定须知_简体中文" path="bookingNotice_zh">
                <n-input type="textarea" placeholder="预定须知" v-model:value="formValue.bookingNotice_zh" @blur="translate('bookingNotice','zh',formValue.bookingNotice_zh)" style="width: 400px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="预定须知_English" path="bookingNotice_en">
                <n-input type="textarea" placeholder="预定须知" v-model:value="formValue.bookingNotice_en" @blur="translate('bookingNotice','en',formValue.bookingNotice_en)" style="width: 400px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="预定须知_日本语" path="bookingNotice_ja">
                <n-input type="textarea" placeholder="预定须知" v-model:value="formValue.bookingNotice_ja" style="width: 400px"/>
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-divider title-placement="left">
                价格说明
              </n-divider>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="价格说明_简体中文" path="priceDesc_zh">
                <n-input type="textarea" placeholder="价格说明" v-model:value="formValue.priceDesc_zh" @blur="translate('priceDesc','zh',formValue.priceDesc_zh)" style="width: 400px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="价格说明_English" path="priceDesc_en">
                <n-input type="textarea" placeholder="价格说明" v-model:value="formValue.priceDesc_en" @blur="translate('priceDesc','en',formValue.priceDesc_en)" style="width: 400px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="价格说明_日本语" path="priceDesc_ja">
                <n-input type="textarea" placeholder="价格说明" v-model:value="formValue.priceDesc_ja" style="width: 400px"/>
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
  const group = ref('foodothersetting');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const formBtnLoading = ref(false);
  const settingStore = useProjectSettingStore();

  const formValue = ref({
    sortType: 1,
    bookingNotice_zh: '',
    bookingNotice_en: '',
    bookingNotice_ja: '',
    bookingNotice_ko: '',
    bookingNotice_zh_CN: '',
    priceDesc_zh: '',
    priceDesc_en: '',
    priceDesc_ja: '',
    priceDesc_ko: '',
    priceDesc_zh_CN: '',
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
    }else if(type == 'priceDesc'){
      if(lang == 'zh'){
        // 简体中文
        formValue.value.priceDesc_zh_CN = ''
      }else if(lang == 'en'){
        // 韩语
        formValue.value.priceDesc_ko = ''
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
        }else if(type == 'priceDesc'){
          if(lang == 'zh'){
            // 简体中文
            formValue.value.priceDesc_zh_CN = _res.zh_CN
          }else if(lang == 'en'){
            // 韩语
            formValue.value.priceDesc_ko = _res.ko
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
