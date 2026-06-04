<template>
  <div class="member-admin-page">
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '25px 20px 0',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">会员等级详情</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="level-detail-div">
            <div class="level-detail-div-title">
              <div></div>
              基础信息
            </div>
            <n-grid y-gap="25" :cols="2">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">等级名称</div>
                  <div class="level-detail-div-item-content">{{ formValue.levelName }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">需要达到经验值</div>
                  <div class="level-detail-div-item-content">{{ formValue.exp }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">等级徽章</div>
                  <div class="level-detail-div-item-content">
                    <n-image style="width: 30px" :src="formValue.levelBadge"/>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">等级卡片</div>
                  <div class="level-detail-div-item-content">
                    <n-image style="width: 150px" :src="formValue.levelCard"/>
                  </div>
                </div>
              </n-gi>
              <n-gi span="2">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">字体颜色</div>
                  <div class="level-detail-div-item-content">
                    <div class="color" :style="{backgroundColor:formValue.wordColor}"></div>
                    {{ formValue.wordColor }}
                  </div>
                </div>
              </n-gi>
              <n-gi span="2">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">会员等级说明</div>
                  <div class="level-detail-div-item-content" style="line-height: 24px">{{ formValue.desc ? formValue.desc : '--' }}</div>
                </div>
              </n-gi>
            </n-grid>
          </div>
          <div class="level-detail-div">
            <div class="level-detail-div-title">
              <div></div>
              权益
            </div>
            <n-grid y-gap="25" :cols="2">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">酒店场景获取积分倍率</div>
                  <div class="level-detail-div-item-content">{{ formValue.hotelGetRate }} 倍</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">餐饮场景获取积分倍率</div>
                  <div class="level-detail-div-item-content">{{ formValue.foodGetRate }} 倍</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">按摩场景获取积分倍率</div>
                  <div class="level-detail-div-item-content">{{ formValue.spaGetRate }} 倍</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">接送机/包车场景获取积分倍率</div>
                  <div class="level-detail-div-item-content">{{ formValue.carGetRate }} 倍</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">储物柜场景获取积分倍率</div>
                  <div class="level-detail-div-item-content">{{ formValue.cabinetGetRate }} 倍</div>
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
import { View } from '@/api/pmsMemberLevel';
import { State, newState, options } from './model';
import { adaModalWidth, getOptionLabel, getOptionTag } from '@/utils/hotgo';

const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref(newState(null));
const dialogWidth = computed(() => {
  return adaModalWidth(580);
});
const fileAvatarCSS = computed(() => {
  return {
    '--n-merged-size': `var(--n-avatar-size-override, 80px)`,
    '--n-font-size': `18px`,
  };
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
      .color{
        width: 20px;
        height: 20px;
        margin-right: 8px;
        border-radius: 2px;
      }
    }
  }
}
</style>

