<template>
  <div class="member-admin-page">
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑会员等级场景 #' + formValue.id : '添加会员等级场景'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="场景名称" path="sceneName">
                  <n-input placeholder="请输入场景名称" v-model:value="formValue.sceneName" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="是否允许积分获取" path="isGetOpen">
                  <n-switch unchecked-value="N" checked-value="Y" v-model:value="formValue.isGetOpen"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="是否允许积分抵扣" path="isPayOpen">
                  <n-switch unchecked-value="N" checked-value="Y" v-model:value="formValue.isPayOpen"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="限制金额" path="limitMoney">
                  <n-input-number placeholder="请输入限制金额" v-model:value="formValue.limitMoney" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="获得积分的抵扣比例" path="getRate">
                  <n-input-number placeholder="请输入获得积分的抵扣比例" v-model:value="formValue.getRate" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="能够使用的总金额比例积分" path="payRate">
                  <n-input-number placeholder="请输入能够使用的总金额比例积分" v-model:value="formValue.payRate" />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import { Edit, View } from '@/api/pmsMemberScene';
import { options, State, newState, rules } from './model';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import { adaModalWidth } from '@/utils/hotgo';

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const settingStore = useProjectSettingStore();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(640);
});

function openModal(state: State) {
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

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          closeForm();
          emit('reloadTable');
        });
      });
    } else {
      message.error('请填写完整信息');
    }
    formBtnLoading.value = false;
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

