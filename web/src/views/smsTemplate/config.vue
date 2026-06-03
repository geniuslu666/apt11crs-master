<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">配置项</text>
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
              <n-form-item label="国内短信通道" path="noticeSmsDriveCN">
                <n-select
                  placeholder="国内短信通道"
                  :options="options.config_internal_sms_drive"
                  v-model:value="formValue.noticeSmsDriveCN"
                  style="width: 300px"
                />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="海外短信通道" path="noticeSmsDriveAbroad">
                <n-select
                  placeholder="海外短信通道"
                  :options="options.config_sms_drive"
                  v-model:value="formValue.noticeSmsDriveAbroad"
                  style="width: 300px"
                />
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
  import {Options} from "@/utils/hotgo";
  import {Dicts} from "@/api/dict/dict";

  const group = ref('sms');
  const rules = ref({});
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const formBtnLoading = ref(false);
  const settingStore = useProjectSettingStore();

  const options = ref<Options>({
    config_sms_drive: [],
    config_internal_sms_drive: []
  });

  const formValue = ref({
    noticeSmsDriveCN: '',
    noticeSmsDriveAbroad: '',
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
      }
    });
  }


  async function loadOptions() {
    options.value = await Dicts({
      types: ['config_sms_drive', 'config_internal_sms_drive'],
    });
  }


  async function load() {
    show.value = true;
    await loadOptions();
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          formValue.value = res.list;
        })
        .finally(() => {
          show.value = false;
        });
    });
  }

  onMounted(() => {
    load();
  });

</script>
