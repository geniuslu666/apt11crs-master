<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form ref="formRef" :label-width="160" label-placement="left" label-align="left" :model="formValue" :rules="rules">
        <n-grid cols="1">
          <n-gi>
            <n-divider title-placement="left">
              语言包配置
            </n-divider>
          </n-gi>
          <n-gi>
            <div style="margin-left: 40px">
              <n-alert :show-icon="false" type="info" style="margin-bottom: 20px;">
                <p>当前语言包版本号：{{ formValue.languagePackVersion }}</p>
                <p>当前语言包文件路径：{{ formValue.languagePackFilePath }}<a :href="formValue.languagePackFilePath" target="_blank">点击下载</a></p>
              </n-alert>

              <template v-if="updateFlag">
                <n-form-item label="语言包版本号" path="languagePackVersion">
                  <n-input-group>
                    <n-input v-model:value="formValue.languagePackVersion" :style="{ width: '240px' }"  />
                  </n-input-group>
                </n-form-item>
                <n-form-item label="语言包文件" path="languagePackFilePath">
                  <UploadFile :maxNumber="1" :accept="`.json`"  v-model:value="formValue.languagePackFilePath" />
                </n-form-item>
              </template>

            </div>
          </n-gi>
        </n-grid>
        <div style="text-align: center">
          <n-space justify="center">
            <n-button type="info"  @click="update">{{ updateText }}</n-button>
            <n-button type="primary" :loading="formBtnLoading" @click="formSubmit" v-if="updateFlag">保存更新</n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, ref} from 'vue';
import {useDialog, useMessage} from 'naive-ui';
import {getConfig, updateConfig} from '@/api/sys/config';
import UploadFile from "@/components/Upload/uploadFile.vue";

const formBtnLoading = ref(false);
const group = ref('languagePackSetting');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const dialog = useDialog();
const updateFlag = ref(false);
const updateText = ref("点击修改");
const oldVersion = ref("");
const oldFilePath = ref("");

  const formValue = ref({
    languagePackVersion: "",
    languagePackFilePath: "",
  });

  // const payMode = ref([])

  const rules = {
  /*  basicName: {
      required: true,
      message: '请输入网站名称',
      trigger: 'blur',
    },*/
  };

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
        formBtnLoading.value = false;
        message.error('验证失败，请填写完整信息');
      }
    });
  }
  function update() {
    if (updateFlag.value) {
      updateText.value = "点击修改";
      updateFlag.value = false;
    }else{
      updateText.value = "取消修改";
      updateFlag.value = true;
    }
    return;
  }

  onMounted(() => {
    load();
  });

  function load() {
    show.value = true;
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          oldVersion.value = res.list.languagePackVersion;
          oldFilePath.value = res.list.languagePackFilePath;
          formValue.value = res.list;
        })
        .finally(() => {
          show.value = false;
        });
    });
  }
</script>
