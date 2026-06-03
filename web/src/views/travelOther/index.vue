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
                核销端-联系方式
              </n-divider>
              <n-gi span="1">
                <n-form-item label="联系电话" path="orderRule_ko">
                  <n-input placeholder="联系电话" v-model:value="formValue.contactMobile" style="width: 400px"/>
                </n-form-item>
              </n-gi>
            </n-gi>
            <n-gi>
              <n-divider title-placement="left">
                预定须知
              </n-divider>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="预定须知_简体中文" path="bookingNotice_zh">
                <n-input type="textarea" placeholder="预定须知" v-model:value="formValue.bookingNotice_zh" style="width: 400px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="预定须知_English" path="bookingNotice_en">
                <n-input type="textarea" placeholder="预定须知" v-model:value="formValue.bookingNotice_en" style="width: 400px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="预定须知_日本语" path="bookingNotice_ja">
                <n-input type="textarea" placeholder="预定须知" v-model:value="formValue.bookingNotice_ja" style="width: 400px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="预定须知_한국어" path="bookingNotice_ko">
                <n-input type="textarea" placeholder="预定须知" v-model:value="formValue.bookingNotice_ko" style="width: 400px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="预定须知_繁体中文" path="bookingNotice_ko">
                <n-input type="textarea" placeholder="预定须知" v-model:value="formValue.bookingNotice_zh_CN" style="width: 400px"/>
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-divider title-placement="left">
                活动规则
              </n-divider>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="活动规则_简体中文" path="orderRule_zh">
                <n-input type="textarea" placeholder="活动规则" v-model:value="formValue.orderRule_zh" style="width: 400px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="活动规则_English" path="orderRule_en">
                <n-input type="textarea" placeholder="活动规则" v-model:value="formValue.orderRule_en" style="width: 400px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="活动规则_日本语" path="orderRule_ja">
                <n-input type="textarea" placeholder="活动规则" v-model:value="formValue.orderRule_ja" style="width: 400px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="活动规则_한국어" path="orderRule_ko">
                <n-input type="textarea" placeholder="活动规则" v-model:value="formValue.orderRule_ko" style="width: 400px"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="活动规则_繁体中文" path="orderRule_ko">
                <n-input type="textarea" placeholder="活动规则" v-model:value="formValue.orderRule_zh_CN" style="width: 400px"/>
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
  const group = ref('travelothersetting');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const formBtnLoading = ref(false);
  const settingStore = useProjectSettingStore();

  const formValue = ref({
    bookingNotice_zh: '',
    bookingNotice_en: '',
    bookingNotice_ja: '',
    bookingNotice_ko: '',
    bookingNotice_zh_CN: '',
    orderRule_zh: '',
    orderRule_en: '',
    orderRule_ja: '',
    orderRule_ko: '',
    orderRule_zh_CN: '',
    contactMobile: '',
  });

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
