<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth" :z-index="99">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑渠道 #' + formValue.id : '添加渠道' }}</div>
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
            :rules="rules"
            label-placement="top"
            label-width="auto"
          >
            <n-grid :cols="2" x-gap="17">
              <n-gi>
                <n-form-item label="渠道名" path="name" >
                  <n-input placeholder="请输入渠道名" v-model:value="formValue.name" :style="{ width: '300px' }" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="手机号" path="phone" >
                  <n-input placeholder="请输入手机号" v-model:value="formValue.phone" :style="{ width: '300px' }" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="邮箱" path="email" >
                  <n-input placeholder="请输入邮箱" v-model:value="formValue.email" :style="{ width: '300px' }" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="返佣比例" path="rate">
                  <n-input-group>
                    <n-input-number :min="0" placeholder="请输入返佣比例" v-model:value="formValue.rate" :show-button="false" :style="{ width: '300px' }" />
                    <n-input-group-label>%</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="最低可提现额" path="minWithdrawalAmount">
                  <n-input-group>
                    <n-input-number :min="0" placeholder="请输入最低可提现额" v-model:value="formValue.minWithdrawalAmount" :show-button="false" :style="{ width: '300px' }" />
                    <n-input-group-label>JPY</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="手续费" path="serviceCharge">
                  <n-input-group>
                    <n-input-number :min="0" placeholder="请输入手续费" v-model:value="formValue.serviceCharge" :show-button="false" :style="{ width: '300px' }" />
                    <n-input-group-label>%</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="预计提现时间周期" path="afterDay">
                  <n-input-group>
                    <n-input-number :min="0" placeholder="请输入预计提现时间周期" v-model:value="formValue.afterDay" :show-button="false" :style="{ width: '300px' }" />
                    <n-input-group-label>天</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="状态" path="status">
                  <n-radio-group v-model:value="formValue.status" name="status">
                    <n-radio-button
                      v-for="status in options.sys_normal_disable"
                      :key="status.value"
                      :value="status.value"
                      :label="status.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="备注" path="remark">
                  <Editor id="remark" v-model:modelValue="formValue.remark" />
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
import { Edit, View } from '@/api/pmsChannel';
import { options, State, newState, rules } from './model';
import Editor from '@/components/Editor/editor.vue';
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
          formBtnLoading.value = false;
          closeForm();
          emit('reloadTable');
        });
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      message.error('请填写完整信息');
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


