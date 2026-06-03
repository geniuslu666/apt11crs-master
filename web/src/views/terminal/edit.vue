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
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑终端 #' + formValue.id : '添加终端' }}</div>
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
            :rules="rules"
            :model="formValue"
            label-placement="top"
            label-width="auto"
          >
            <n-grid :cols="1">
              <n-gi>
                <n-form-item label="终端名称" path="terminalName" :show-require-mark="true">
                  <n-input placeholder="请输入终端名称" v-model:value="formValue.terminalName"  style="width: 300px"/>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="终端类型" path="terminalType">
                  <a-radio-group v-model:value="formValue.terminalType" name="radioGroup" @update:value="changeType">
                    <a-radio value="VERIFY_PRINTER">核销打印机</a-radio>
                    <a-radio value="HAND_TERMINAL">手持终端</a-radio>
                  </a-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.terminalType == 'VERIFY_PRINTER'">
                <n-form-item label="品牌型号" path="brandModel">
                  <n-select
                    v-model:value="formValue.brandModel"
                    :options="options.verify_brand_model"
                    style="width: 300px"
                  />
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.terminalType == 'HAND_TERMINAL'">
                <n-form-item label="品牌型号" path="brandModel">
                  <n-select
                    v-model:value="formValue.brandModel"
                    :options="options.hand_brand_model"
                    style="width: 300px"
                    label-field="label"
                    value-field="value"
                  />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="终端编号" path="sn" :show-require-mark="true">
                  <n-input placeholder="请输入终端编号" v-model:value="formValue.sn"  style="width: 300px"/>
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
import { ref, computed } from 'vue';
import { Edit, View } from '@/api/terminal';
import { options, State, newState, rules } from './model';
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

async function openModal(state: State) {
  showModal.value = true;

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);
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

function changeType(e) {
  console.log(e);
  if(e == 'VERIFY_PRINTER'){
    formValue.value.brandModel = 'XPYUN503'
  }else{
    formValue.value.brandModel = 'SHANGMI_V3'
  }
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


