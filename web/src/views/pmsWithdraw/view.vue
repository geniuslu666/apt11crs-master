<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '25px 0 0',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">提现申请表详情</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="level-detail-div" style="padding: 0 20px;">
            <n-grid y-gap="25" :cols="2">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">提现单号</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.withdrawSn }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">员工ID</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.staffId }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">员工名称</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.pmsStaffName }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">提现状态</div>
                  <div class="level-detail-div-item-content">
                    <div class="level-detail-div-item-content-tag1 success" v-if="formValue.withdrawStatus == 'WAIT'">待审核</div>
                    <div class="level-detail-div-item-content-tag1 default" v-else-if="formValue.withdrawStatus == 'SUCCESS'">已同意</div>
                    <div class="level-detail-div-item-content-tag1 error" v-else-if="formValue.withdrawStatus == 'FAIL'">已拒绝</div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">转账状态</div>
                  <div class="level-detail-div-item-content">
                    <div class="level-detail-div-item-content-tag1 success" v-if="formValue.transfer == 1">待转账</div>
                    <div class="level-detail-div-item-content-tag1 default" v-else-if="formValue.transfer == 2">已转账</div>
                    <div class="level-detail-div-item-content-tag1" v-else>--</div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">提现金额</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.withdrawAmount }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">到账金额</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.arrivalAmount }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">提现手续费比例</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.serviceCharge }}%
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">审核时间</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.applyAt }}
                  </div>
                </div>
              </n-gi>
              <n-gi :span="2">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">审核备注</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.applyRemark ? formValue.applyRemark :  '--' }}
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </div>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue';
import { useMessage } from 'naive-ui';
import { View } from '@/api/pmsWithdraw';
import { State, newState } from './model';
import { adaModalWidth } from '@/utils/hotgo';

const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref(newState(null));
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});
const fileAvatarCSS = computed(() => {
  return {
    '--n-merged-size': `var(--n-avatar-size-override, 80px)`,
    '--n-font-size': `18px`,
  };
});

//下载
function download(url: string) {
  window.open(url);
}

function openModal(state: State) {
  showModal.value = true;
  loading.value = true;
  View({ id: state.id })
    .then((res) => {
      formValue.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
}

defineExpose({
  openModal,
});
</script>

<style lang="less" scoped>

.level-detail-div{
  margin-bottom: 25px;
  &-title{
    display: flex;
    align-items: center;
    font-weight: 500;
    font-size: 16px;
    color: #3D3D3D;
    line-height: 22px;
    margin-bottom: 25px;
    div{
      margin-right: 5px;
      width: 6px;
      height: 15px;
      background: #053DC8;
    }
  }
  &-item{
    &-title{
      font-weight: 400;
      font-size: 12px;
      color: #707070;
      line-height: 17px;
      margin-bottom: 8px;
    }
    &-content{
      font-weight: 500;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      display: flex;
      align-items: center;
      a{
        font-weight: 400;
        font-size: 14px;
        color: #1577FF;
        line-height: 22px;
        margin-left: 6px;
        cursor: pointer;
      }
      .color{
        width: 20px;
        height: 20px;
        margin-right: 8px;
        border-radius: 2px;
      }
      &-tag1{
        padding: 0 5px;
        height: 22px;
        line-height: 22px;
        text-align: center;
        font-size: 14px;
        border-radius: 2px;
        font-weight: 400;
        &.success{
          color: #26A763;
          background: #E3F4EB;
        }
        &.warning{
          color: #EFA020;
          background: #FFECCE;
        }
        &.primary{
          color: #3F9EFF;
          background: #ECF5FF;
        }
        &.error{
          color: #F56C6C;
          background: #FEF0F0;
        }
        &.default{
          background: #F4F4F5;
          color: #919399;
        }
      }
    }
  }
}
</style>


