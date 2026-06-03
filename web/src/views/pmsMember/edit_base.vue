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
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ baseTitle }}</div>
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
            <n-grid :cols="1">
              <n-gi v-if="formValue.type == 'group'">
                <n-form-item label="会员分组" path="group">
                  <n-select
                    placeholder="请选择会员分组"
                    v-model:value="formValue.groupId"
                    :options="groupList"
                    label-field="memberGroup"
                    value-field="id"
                    clearable
                    filterable
                    style="width: 300px"
                  />
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.type == 'level'">
                <n-form-item label="会员等级" path="level">
                  <n-select
                    placeholder="请选择会员等级"
                    v-model:value="formValue.level"
                    :options="levelList"
                    label-field="levelName"
                    value-field="id"
                    clearable
                    filterable
                    style="width: 300px"
                  />
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.type == 'lastName'">
                <n-form-item label="会员名" path="lastName">
                  <n-input placeholder="请输入会员名" v-model:value="formValue.lastName" style="width: 300px"/>
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.type == 'firstName'">
                <n-form-item label="会员姓" path="firstName">
                  <n-input placeholder="请输入会员姓" v-model:value="formValue.firstName" style="width: 300px"/>
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.type == 'phone'">
                <n-form-item label="区号" path="phoneArea">
                  <n-input placeholder="请输入区号" v-model:value="formValue.phoneArea" style="width: 300px"/>
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.type == 'phone'">
                <n-form-item label="手机号" path="phone">
                  <n-input placeholder="请输入手机号" v-model:value="formValue.phone" style="width: 300px"/>
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.type == 'mail'">
                <n-form-item label="邮箱" path="mail">
                  <n-input placeholder="请输入邮箱" v-model:value="formValue.mail" style="width: 300px"/>
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
import {computed, ref} from 'vue';
  import { BaseEdit } from '@/api/pmsMember';
  import { useMessage, FormItemRule } from 'naive-ui';
  import {All} from "@/api/pmsMemberLevel/index";
  import {GroupAll} from "@/api/pmsMemberGroup/index";
import {adaModalWidth} from "@/utils/hotgo";

  const baseTitle = ref('')
  const levelList = ref([])
  const groupList = ref([])
  const emit = defineEmits(['reloadInfo']);
  const message = useMessage();
  const loading = ref(false);
  const showModal = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(650);
  });
  const formValue = ref({
    id: '',
    type: '',
    level: '',
    groupId: '',
    firstName: '',
    lastName: '',
    phoneArea: '',
    phone: '',
    mail: ''
  })
  const formRef = ref<any>({});
  const rules = ref({});
  const formBtnLoading = ref(false);

  function openModal(id, baseType, value) {
    formValue.value.id = id
    formValue.value.type = baseType
    formValue.value.level = value.level
    formValue.value.groupId = value.groupId > 0 ? value.groupId : ''
    formValue.value.firstName = value.firstName
    formValue.value.lastName = value.lastName
    formValue.value.phoneArea = value.phoneArea
    formValue.value.phone = value.phone
    formValue.value.mail = value.mail

    if(baseType == 'level'){
      loading.value = true;
      All({}).then((res) => {
        levelList.value = res.list;
        baseTitle.value = '编辑会员等级'
        loading.value = false;
      });
    }else if(baseType == 'group'){
      loading.value = true;
      GroupAll({}).then((res) => {
        groupList.value = res.list;
        baseTitle.value = '编辑会员分组'
        loading.value = false;
      });
    }else if(baseType == 'firstName'){
      baseTitle.value = '编辑会员姓'
    }else if(baseType == 'lastName'){
      baseTitle.value = '编辑会员名'
    }else if(baseType == 'phone'){
      baseTitle.value = '编辑手机号'
    }else if(baseType == 'mail'){
      baseTitle.value = '编辑邮箱'
      rules.value = {
        mail: {
          required: true,
          trigger: ['blur', 'input'],
          validator(rule: FormItemRule, value: string) {
            if (!value) {
              return new Error('邮箱必填')
            }
            else if (!/^\w+(\.)?(\w+)?@[0-9a-z]+(\.[a-z]+){1,3}$/.test(value)) {
              return new Error('邮箱格式不正确')
            }
            return true
          },
        },
      };
    }
    showModal.value = true;
  }

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        BaseEdit(formValue.value).then((_res) => {
          message.success('操作成功');
          setTimeout(() => {
            closeForm();
            formBtnLoading.value = false;
            emit('reloadInfo');
          });
        }).catch((err) => {
          formBtnLoading.value = false;
        });
      } else {
        message.error('验证错误');
        formBtnLoading.value = false;
      }
    });
  }

  function closeForm() {
    showModal.value = false;
    loading.value = false;
    formValue.value = {
      id: '',
      type: '',
      level: '',
      firstName: '',
      lastName: '',
      phoneArea: '',
      phone: '',
      mail: ''
    }
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>
