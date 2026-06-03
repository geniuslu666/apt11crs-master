<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑会员信息 #' + formValue.id : '添加会员信息'"
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
            :label-width="160"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="会员号" path="memberNo">
                  <n-input placeholder="请输入会员号" v-model:value="formValue.memberNo" />
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
                <n-form-item label="全名" path="fullName">
                  <n-input placeholder="请输入全名" v-model:value="formValue.fullName" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="区号" path="phone">
                  <n-input placeholder="请输入手机区号" v-model:value="formValue.phoneArea" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="手机号" path="phone">
                  <n-input placeholder="请输入手机号" v-model:value="formValue.phone" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="邮箱" path="mail">
                  <n-input placeholder="请输入邮箱" v-model:value="formValue.mail" />
                </n-form-item>
              </n-gi>

              <n-gi span="1">
                <n-form-item label="生日" path="birthday">
                  <DatePicker v-model:formValue="formValue.birthday" type="date" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="积分" path="balance">
                  {{ formValue.balance }}
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="推荐人" path="referrer">
                  {{ formValue.referrer }}
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="注册来源  IOS,Andriod,h5" path="source">
                  <!-- <n-input placeholder="请输入注册来源  IOS,Andriod,h5" v-model:value="formValue.source" /> -->
                  {{ formValue.source }}
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="上次登录IP" path="lastLoginIp">
                  <!-- <n-input placeholder="请输入上次登录IP" v-model:value="formValue.lastLoginIp" /> -->
                  {{ formValue.lastLoginIp }}
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="上次登录时间" path="lastLogin">
                  <!-- <DatePicker v-model:formValue="formValue.lastLogin" type="datetime" /> -->
                  {{ formValue.lastLogin }}
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="注册IP" path="registerIp">
                  <!-- <n-input placeholder="请输入注册IP" v-model:value="formValue.registerIp" /> -->
                  {{ formValue.registerIp }}
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="注册时间" path="registerTime">
                  <!-- <DatePicker v-model:formValue="formValue.registerTime" type="datetime" /> -->
                  {{ formValue.registerTime }}
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="地址" path="address">
                  <n-input placeholder="请输入地址" v-model:value="formValue.address" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="国籍" path="nationality">
                  <!-- <n-input placeholder="请输入国籍" v-model:value="formValue.nationality" /> -->
                  {{formValue.nationality}}
                </n-form-item>
              </n-gi>
           
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm"> 取消 </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm"> 确定 </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref, computed } from 'vue';
  import { Edit, View } from '@/api/pmsMember';
  import { State, newState, rules } from './model';
  import DatePicker from '@/components/DatePicker/datePicker.vue';
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
