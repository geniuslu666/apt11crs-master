<template>
  <div class="member-admin-page">
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '25px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">选择启用日期</div>
        </template>
        <template #footer>
          <n-button @click="closeForm" style="width: 70px;height: 35px;margin-right: 10px">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm" style="width: 70px;height: 35px;">
            发放
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <n-form :model="formValue" ref="formRef"  label-placement="top"
                  label-width="auto"
                  require-mark-placement="right-hanging"
          >
            <n-grid :cols="1">
              <n-gi>
                <n-form-item label="启用时间" path="startTime" >
                  <n-date-picker format="yyyy-MM-dd HH:mm:ss" v-model:formatted-value="formValue.startTime" type="datetime" clearable />
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
import {adaModalWidth} from "@/utils/hotgo";
import {SendMemberCoupon} from '@/api/thCoupon';
import {useDialog, useMessage} from "naive-ui";

const emit = defineEmits(['closeDrawer']);
const dialog = useDialog();
const message = useMessage();
const formBtnLoading = ref(false);
  const loading = ref(false);
  const showModal = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(650);
  });
  const formValue = ref({
    startTime: null
  })

const couponId = ref(0);
  const couponMemberId = ref(0);

  function confirmForm(e){
    if(!formValue.value.startTime){
      message.error('请选择启用时间');
      return false;
    }
    dialog.warning({
      title: '操作确认',
      content: '此操作会将礼品券直接发放至对应会员礼品券账户，是否确认操作?',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        e.preventDefault();
        formBtnLoading.value = true;
        SendMemberCoupon({
          memberId: couponMemberId.value,
          couponId: couponId.value,
          startTime: formValue.value.startTime
        }).then((_res) => {
          message.success('发放成功');
          formBtnLoading.value = false;
          // reloadTable();
          setTimeout(() => {
            closeForm();
            emit('closeDrawer');
          }, 1100);
        }).catch((err) => {
          formBtnLoading.value = false;
        });
      },
    });
  }

  function openModal(memberId, chooseCouponId) {
    formValue.value.startTime = null;
    couponId.value = 0;
    couponMemberId.value = 0;
    showModal.value = true;
    couponId.value = chooseCouponId;
    couponMemberId.value = memberId
  }

function closeForm() {
  showModal.value = false;
}

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>
