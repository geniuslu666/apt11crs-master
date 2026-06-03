<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard"
        size="small"
        :segmented="{ content: true }"
      >
        <n-form
          ref="formRef"
          :model="formValue"
          :rules="rules"
          :label-placement="settingStore.isMobile ? 'top' : 'left'"
          :label-width="200"
          class="py-4"
        >
          <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen" y-gap="10">
            <n-gi span="1">
              <n-form-item label="开发者ID" path="clientId">
                <n-input v-model:value="formValue.clientId" style="width: 300px"
                         placeholder="请输入开发者ID"/>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="开发者秘钥" path="clientSecret">
                <n-input v-model:value="formValue.clientSecret" style="width: 300px"
                         placeholder="请输入开发者SECRET"/>
              </n-form-item>
            </n-gi>
          </n-grid>
          <div style="text-align: center">
            <n-space justify="center">
              <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
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
import {ref, onMounted, watch} from 'vue';
import { BrandEdit, BrandView } from '@/api/terminal';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';

interface Props {
  brandId?: string;
}

const props = withDefaults(defineProps<Props>(), {
  brandId: '0',
});
const emit = defineEmits(['reloadTable']);
const message = useMessage();
const settingStore = useProjectSettingStore();
const loading = ref(false);
const formValue = ref({
  id: '',
  brandModel: '',
  clientId: '',
  clientSecret: ''
});
const formRef = ref<any>({});
const formBtnLoading = ref(false);

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      BrandEdit(formValue.value).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          emit('reloadTable');
        });
      });
    } else {
      message.error('请填写完整信息');
    }
    formBtnLoading.value = false;
  });
}


onMounted(() => {
  load();
});

watch(
  () => props.brandId,
  (value) => {
    loading.value = true;
    BrandView({ id: value })
      .then((res) => {
        formValue.value = res;
        console.log('----------------------------',formValue.value)
      })
      .finally(() => {
        loading.value = false;
      });
  }
);

function load() {
  console.log('props.brandId',props.brandId)
  // 编辑
  // loading.value = true;
  // BrandView({ id: props.brandId })
  //   .then((res) => {
  //     formValue.value = res;
  //     console.log('----------------------------',formValue.value)
  //   })
  //   .finally(() => {
  //     loading.value = false;
  //   });
}

</script>

<style lang="less"></style>


