<template>
  <div class="member-admin-page">
    <n-spin :show="show" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
      >
        <n-form :label-width="100" :model="submitFormValue" :rules="rules" ref="formRef" label-placement="left">

          <n-form-item label="语言" path="language">
            <n-radio-group v-model:value="submitFormValue.language" name="language"  @update:value="handleUpdateLanguageValue">
              <n-radio-button
                v-for="language in options.language"
                :key="language.value"
                :value="language.value"
                :label="language.label"
              />
            </n-radio-group>
          </n-form-item>

          <template v-for="language in options.language">
            <n-form-item :key="'rule_'+language.value" label="等级说明" :path="'rule_'+language.value" style="width: 100%;max-width: 1200px"  v-if="submitFormValue.language == language.value" >
              <Editor :id="'rule_id_'+language.value" v-model:modelValue="submitFormValue.rule"/>
            </n-form-item>
          </template>

          <div>
            <n-space justify="center">
              <n-button type="primary" @click="formSubmit">保存更新</n-button>
            </n-space>
          </div>
        </n-form>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { useMessage } from 'naive-ui';
import { getConfig, updateConfig } from '@/api/sys/config';
import Editor from '@/components/Editor/editor.vue';
import {options, loadOptions} from "@/views/pmsMemberLevel/model";
import {useTabsViewStore} from "@/store/modules/tabsView";

const tabsViewStore = useTabsViewStore();

const group = ref('memberlevel');
const show = ref(false);
const rules = {};
const formRef: any = ref(null);
const message = useMessage();
const formValue = ref({
  rule: '',
});


const submitFormValue = ref({
  language: 'zh',
  rule: '',
})

function formSubmit() {
  console.log('formValue.value', formValue.value)
  console.log('formValue.submitFormValue', submitFormValue.value)


  formRef.value.validate((errors) => {
    if (!errors) {
      formValue.value['rule_'+submitFormValue.value.language] = submitFormValue.value.rule
      updateConfig({ group: group.value, list: formValue.value })
        .then((_res) => {
          message.success('更新成功');
          tabsViewStore.closeSignal('1');
        })
        .catch((error) => {
          message.error(error.toString());
        });
    } else {
      message.error('验证失败，请填写完整信息');
    }
  });
}

function handleUpdateLanguageValue(value) {

  submitFormValue.value.rule = formValue.value['rule_'+value]

  console.log(submitFormValue.value.rule)
}

onMounted(() => {
  loadOptions();
  load();
});

function load() {
  show.value = true;
  new Promise((_resolve, _reject) => {
    getConfig({ group: group.value })
      .then((res) => {
        show.value = false;
        if(res.list){
          formValue.value = res.list;
          submitFormValue.value.rule = res.list['rule_'+submitFormValue.value.language]
        }
      })
      .catch((error) => {
        show.value = false;
        message.error(error.toString());
      });
  });
}
</script>
