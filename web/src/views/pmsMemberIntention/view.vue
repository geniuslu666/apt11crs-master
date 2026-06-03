<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '25px 0 0',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">会员意向表详情</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="level-detail-div" style="padding: 0 20px;">
            <n-grid y-gap="25" :cols="2">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">会员ID</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.memberId }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">会员编号</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.member?.memberNo }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">语言</div>
                  <div class="level-detail-div-item-content">
                    <div class="level-detail-div-item-content-tag1 error" v-if="formValue.language == '简体中文'">{{ formValue.language }}</div>
                    <div class="level-detail-div-item-content-tag1 warning" v-else-if="formValue.language == 'English'">{{ formValue.language }}</div>
                    <div class="level-detail-div-item-content-tag1 primary" v-else-if="formValue.language == '日本語'">{{ formValue.language }}</div>
                    <div class="level-detail-div-item-content-tag1 success" v-else-if="formValue.language == '한국어'">{{ formValue.language }}</div>
                    <div class="level-detail-div-item-content-tag1 default" v-else-if="formValue.language == '繁体中文'">{{ formValue.language }}</div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">国家/地区</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.country }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">邮箱</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.mail }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">手机号</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.phoneArea }} {{ formValue.phone }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">LINE ID</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.lineId }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">微信号</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.wechatNo }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">创建时间</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.createAt }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">修改时间</div>
                  <div class="level-detail-div-item-content">
                    {{ formValue.updateAt }}
                  </div>
                </div>
              </n-gi>
              <n-gi span="2">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">备注</div>
                  <div class="level-detail-div-item-content">
                    <div v-html="formValue.remark" v-if="formValue.remark"></div>
                    <div v-else>无</div>
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
import {computed, ref} from 'vue';
import { View } from '@/api/pmsMemberIntention';
import {State, newState} from './model';
import { adaModalWidth } from '@/utils/hotgo';


const loading = ref(false);
const showModal = ref(false);
const formValue = ref(newState(null));
const dialogWidth = computed(() => {
  return adaModalWidth(580);
});

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
          color: #434447;
          background: #EEEEEE;
        }
      }
    }
  }
}
</style>


