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
              <n-gi v-if="formValue.type == 'name'">
                <n-form-item label="姓名" path="name" >
                  <n-input placeholder="请输入姓名" v-model:value="formValue.name" :style="{ width: '300px' }" />
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.type == 'phone'">
                <n-form-item label="手机号" path="phone">
                  <n-input placeholder="请输入手机号" v-model:value="formValue.phone" :style="{ width: '300px' }"/>
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.type == 'email'">
                <n-form-item label="邮箱" path="email">
                  <n-input placeholder="请输入邮箱" v-model:value="formValue.email" :style="{ width: '300px' }"/>
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.type == 'rate'">
                <n-form-item label="返佣比例" path="rate">
                  <n-input-group>
                    <n-input-number :min="0" placeholder="请输入返佣比例" v-model:value="formValue.rate" :show-button="false" :style="{ width: '300px' }" />
                    <n-input-group-label>%</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.type == 'minWithdrawalAmount'">
                <n-form-item label="最低可提现额" path="minWithdrawalAmount">
                  <n-input-group>
                    <n-input-number :min="0" placeholder="请输入最低可提现额" v-model:value="formValue.minWithdrawalAmount" :show-button="false" :style="{ width: '300px' }" />
                    <n-input-group-label>JPY</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.type == 'serviceCharge'">
                <n-form-item label="手续费" path="serviceCharge">
                  <n-input-group>
                    <n-input-number :min="0" placeholder="请输入手续费" v-model:value="formValue.serviceCharge" :show-button="false" :style="{ width: '300px' }" />
                    <n-input-group-label>%</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.type == 'afterDay'">
                <n-form-item label="预计提现时间周期" path="afterDay">
                  <n-input-group>
                    <n-input-number :min="0" placeholder="请输入预计提现时间周期" v-model:value="formValue.afterDay" :show-button="false" :style="{ width: '300px' }" />
                    <n-input-group-label>天</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi v-if="formValue.type == 'status'">
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
            </n-grid>
          </n-form>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {computed, ref} from 'vue';
  import { Edit } from '@/api/pmsChannel';
  import { useMessage } from 'naive-ui';
  import {validate} from "@/utils/validateUtil";
  import {options} from "@/views/pmsChannel/model";
import {adaModalWidth} from "@/utils/hotgo";

  const baseTitle = ref('')
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
    name: '',
    phone: '',
    email: '',
    rate: 0,
    minWithdrawalAmount: 0,
    serviceCharge: 0,
    afterDay: 0,
    status: 0,
  })
  const formRef = ref<any>({});
  const rules = ref({});
  const formBtnLoading = ref(false);

  function openModal(id, baseType, value) {
    formValue.value.id = id
    formValue.value.type = baseType
    formValue.value.name = value.name
    formValue.value.phone = value.phone
    formValue.value.email = value.email
    formValue.value.rate = value.rate
    formValue.value.minWithdrawalAmount = value.minWithdrawalAmount
    formValue.value.serviceCharge = value.serviceCharge
    formValue.value.afterDay = value.afterDay
    formValue.value.status = value.status

    if(baseType == 'name'){
      baseTitle.value = '编辑姓名'
    }else if(baseType == 'phone'){
      baseTitle.value = '编辑手机号'
      rules.value = {
        phone: {
          trigger: ['blur', 'input'],
          validator: validate.phone,
        },
      };
    }else if(baseType == 'email'){
      baseTitle.value = '编辑邮箱'
      rules.value = {
        email: {
          trigger: ['blur', 'input'],
          validator: validate.email,
        },
      };
    }else if(baseType == 'rate'){
      baseTitle.value = '编辑返佣比例'
      rules.value = {
        rate: {
          trigger: ['blur', 'input'],
          validator: validate.percentage
        },
      };
    }else if(baseType == 'minWithdrawalAmount'){
      baseTitle.value = '编辑最低可提现额'
    }else if(baseType == 'serviceCharge'){
      baseTitle.value = '编辑提现手续费'
    }else if(baseType == 'afterDay'){
      baseTitle.value = '编辑提现周期'
    }else if(baseType == 'status'){
      baseTitle.value = '编辑状态'
    }
    showModal.value = true;
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
      name: '',
      phone: '',
      email: '',
      rate: 0,
      minWithdrawalAmount: 0,
      serviceCharge: 0,
      afterDay: 0,
      status: 0,
    }
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>
