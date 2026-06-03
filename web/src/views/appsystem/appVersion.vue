<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="80" :model="formValue" :rules="rules" ref="formRef">
        <n-form-item label="机型" path="phoneModelArr">
          <n-dynamic-tags v-model:value="phoneModelArr" :on-update:value="handleUpdateVersion"/>
        </n-form-item>
        <n-tabs type="line" animated v-model:value="tabValue">

          <n-tab-pane v-for="versionItem in appVersionArr" :key="versionItem.name" :name="versionItem.name">
            <n-form-item label="app版本号" path="appversion">
              <n-input v-model:value="versionItem.info.app_version" placeholder="请输入app版本号" />
            </n-form-item>
            <n-form-item label="app隐私协议" path="app_privacy">
              <n-input v-model:value="versionItem.info.app_privacy" placeholder="请输入app隐私协议" />
            </n-form-item>
          </n-tab-pane>
        </n-tabs>
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
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';

  const group = ref('app');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const phoneModelArr = ref([])
  const appVersionArr = ref([
    {
      name: 'ios',
      info: {
        app_version: '1.0.0',
        app_privacy: '',
      }
    }
  ])
  const tabValue = ref('ios')
  const formValue = ref({
    appversion: '',
  });
  const rules = {
    // basicName: {
    //   required: true,
    //   message: '请输入网站名称',
    //   trigger: 'blur',
    // },
  };

  function formSubmit() {

    formValue.value.appversion = JSON.stringify(appVersionArr.value)

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

  function handleUpdateVersion(values: string[]) {
    phoneModelArr.value = values;
    let values_length = values.length;
    let model_length = appVersionArr.value.length;
    if(model_length < values_length){
      appVersionArr.value.push({
        name: values[values_length - 1],
        info: {
          app_version: '1.0.0',
          app_privacy: '',
        }
      })
    }else if(model_length > values_length){
      let diffIndexes = [];
      appVersionArr.value.forEach((item, index) => {
        if (values.indexOf(item.name) === -1) {
          diffIndexes.push(index);
        }
      });
      // console.log("diffIndexes",diffIndexes)
      diffIndexes.forEach(index => {
        appVersionArr.value.splice(index, 1);
      });
    }

  }

  onMounted(() => {
    load();
  });

  function load() {
    show.value = true;
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          formValue.value.appversion = res.list.appversion;

          let arr = JSON.parse(formValue.value.appversion);
          tabValue.value =  arr[0]['name'];
          appVersionArr.value = arr;

          phoneModelArr.value = arr.map(item => String(item.name));

        })
        .finally(() => {
          show.value = false;
        });
    });
  }
</script>
