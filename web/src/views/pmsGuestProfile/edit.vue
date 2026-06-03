<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑客户档案 #' + formValue.id : '添加客户档案'"
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
                <n-form-item label="客户档案ID" path="uid">
                  <n-input placeholder="请输入客户档案ID" v-model:value="formValue.uid" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="名" path="firstName">
                  <n-input placeholder="请输入名" v-model:value="formValue.firstName" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="姓" path="lastName">
                  <n-input placeholder="请输入姓" v-model:value="formValue.lastName" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="名的假名" path="firstNameKana">
                  <n-input placeholder="请输入名的假名" v-model:value="formValue.firstNameKana" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="姓的假名" path="lastNameKana">
                  <n-input placeholder="请输入姓的假名" v-model:value="formValue.lastNameKana" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="全名" path="fullName">
                  <n-input placeholder="请输入全名" v-model:value="formValue.fullName" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="语言" path="language">
                  <n-input placeholder="请输入语言" v-model:value="formValue.language" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="电子邮件" path="email">
                  <n-input placeholder="请输入电子邮件" v-model:value="formValue.email" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="电话" path="phone">
                  <n-input placeholder="请输入电话" v-model:value="formValue.phone" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="国籍" path="nationality">
                  <n-input placeholder="请输入国籍" v-model:value="formValue.nationality" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="地址" path="address">
                  <n-input placeholder="请输入地址" v-model:value="formValue.address" />
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
  import { Edit, View } from '@/api/pmsGuestProfile';
  import { State, newState, rules } from './model';
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
    return adaModalWidth(840);
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