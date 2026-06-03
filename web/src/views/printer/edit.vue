<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑打印机 #' + formValue.id : '添加打印机' }}</div>
        </template>
        <template #footer>
          <n-button @click="closeForm" style="width: 70px;height: 35px;margin-right: 10px">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm" style="width: 70px;height: 35px;">
            保存
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            label-placement="top"
            label-width="auto"
          >
            <n-grid :cols="2" x-gap="17">
              <n-gi>
                <n-form-item label="打印机名称" path="printerName" :show-require-mark="true">
                  <n-input placeholder="请输入打印机名称" v-model:value="formValue.printerName" :style="{ width: '300px' }" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="打印机类型" path="printer_type">
                  <n-select
                    v-model:value="formValue.printerType"
                    :options="languageOptions"
                    style="width: 300px"
                  />
                </n-form-item>
              </n-gi>
              <template v-if="formValue.printerType == 'YILINK'">
                <n-gi>
                  <n-form-item label="第三方应用ID" path="clientId" :show-require-mark="true">
                    <n-input placeholder="请输入易联云第三方应用ID" v-model:value="formValue.clientId"  style="width: 300px"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="第三方应用秘钥" path="clientSecret" :show-require-mark="true">
                    <n-input placeholder="请输入易联云第三方应用秘钥" v-model:value="formValue.clientSecret"  style="width: 300px"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="打印机终端号" path="machineCode" :show-require-mark="true">
                    <n-input placeholder="请输入打印机终端号" v-model:value="formValue.machineCode"  style="width: 300px"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="终端秘钥" path="machineKey" :show-require-mark="true">
                    <n-input placeholder="请输入打印机终端秘钥" v-model:value="formValue.machineKey"  style="width: 300px"/>
                  </n-form-item>
                </n-gi>
              </template>
              <template v-if="formValue.printerType == 'XPYUN503'">
                <n-gi>
                  <n-form-item label="开发者ID" path="clientId" :show-require-mark="true">
                    <n-input v-model:value="formValue.clientId" style="width: 300px"
                             placeholder="请输入开发者ID"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="开发者秘钥" path="clientSecret" :show-require-mark="true">
                    <n-input v-model:value="formValue.clientSecret" style="width: 300px"
                             placeholder="请输入开发者SECRET"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="终端编号" path="machineCode" :show-require-mark="true">
                    <n-input placeholder="请输入终端编号" v-model:value="formValue.machineCode"  style="width: 300px"/>
                  </n-form-item>
                </n-gi>
              </template>
              <n-gi>
                <n-form-item label="打印联数" path="printTimes">
                  <n-input-group>
                    <n-input-number placeholder="请输入" :show-button="false" :min="0" :precision="0" v-model:value="formValue.printTimes" style="width: 300px" />
                    <n-input-group-label>联</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="排序" path="sort">
                  <n-input-number
                    placeholder="请输入排序"
                    v-model:value="formValue.sort"
                    clearable
                    :show-button="false"
                    style="width: 300px"
                  />
                  <template #feedback>
                    <div style="font-size: 12px">越大越靠前</div>
                  </template>
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {ref, computed, reactive} from 'vue';
import { Edit, View, MaxSort } from '@/api/printer';
import { options, State, newState } from './model';
import { useMessage } from 'naive-ui';
import { adaModalWidth } from '@/utils/hotgo';

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});

const languageOptions = reactive([
  {
    value: 'YILINK',
    label: '易联云',
    disabled: false,
  },
  {
    value: 'XPYUN503',
    label: '芯烨云 ',
    disabled: false,
  },
]);

async function loadMax(){
  let max = await MaxSort()
  formValue.value.sort = max.sort
}

async function openModal(state: State) {
  showModal.value = true;

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);
    await loadMax()
    return;
  }

  // 编辑
  loading.value = true;
  View({ id: state.id })
    .then((res) => {
      formValue.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          formBtnLoading.value = false;
          closeForm();
          emit('reloadTable');
        });
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      message.error('请填写完整信息');
      formBtnLoading.value = false;
    }
  });
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
}

defineExpose({
  openModal,
});
</script>

<style lang="less"></style>


