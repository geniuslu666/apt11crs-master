<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="80" :model="formValue" :rules="rules" ref="formRef">
        <n-form-item label="首页招募广告banner" path="basicLogo">
          <FileChooser v-model:value="formValue.homepageAdvBanner" />
        </n-form-item>
        <div>
          <n-space>
            <n-button type="primary" @click="formSubmit">保存更新</n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import FileChooser from '@/components/FileChooser/index.vue';

  const group = ref('appconfig');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const dialog = useDialog();

  const formValue = ref({
    homepageAdvBanner: '',
  });

  const rules = {

  };

  function formSubmit() {
    formRef.value.validate((errors) => {
      if (!errors) {
        updateConfig({ group: group.value, list: formValue.value }).then((_res) => {
          message.success('更新成功');
          load();
        });
      } else {
        message.error('验证失败，请填写完整信息');
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
          formValue.value = res.list;
        })
        .finally(() => {
          show.value = false;
        });
    });
  }
</script>
