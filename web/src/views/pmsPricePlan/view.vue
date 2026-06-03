<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px 20px 0',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.planName }}</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="level-detail-div">
            <div class="level-detail-div-title">
              <div></div>
              价格计划设置
            </div>
            <n-grid y-gap="25" :cols="2">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">当前状态</div>
                  <div>
                    <div v-if="formValue.pricePlanStatus == 'Y'" class="level-detail-div-item-content1 success">启用中</div>
                    <div v-else class="level-detail-div-item-content1 warning">已停用</div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">取消政策</div>
                  <div class="level-detail-div-item-content">{{ formValue.isCancel == "Y" ? '灵活取消' : '不可取消' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">房价Plan</div>
                  <div class="level-detail-div-item-content">房价比Standard Rate{{ formValue.priceMode == '-' ? '便宜' : '贵' }}{{ formValue.priceStandard == 'PERCENT' ? formValue.planValue + '%' : 'JPY' + formValue.planValue }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">关联房型</div>
                  <div class="level-detail-div-item-content">{{ formValue.roomTypeDetail.name }}</div>
                </div>
              </n-gi>
            </n-grid>
          </div>
          <div class="level-detail-div">
            <div class="level-detail-div-title">
              <div></div>
              价格计划显示名称
            </div>
            <n-grid y-gap="25" :cols="1">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">简体中文名称</div>
                  <div class="level-detail-div-item-content">{{ nameLanguage.zh }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">日文名称</div>
                  <div class="level-detail-div-item-content">{{ nameLanguage.ja }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">英文名称</div>
                  <div class="level-detail-div-item-content">{{ nameLanguage.en }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">韩文名称</div>
                  <div class="level-detail-div-item-content">{{ nameLanguage.ko }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">繁体中文名称</div>
                  <div class="level-detail-div-item-content">{{ nameLanguage.zh_CN }}</div>
                </div>
              </n-gi>
            </n-grid>
          </div>
          <div class="level-detail-div">
            <div class="level-detail-div-title">
              <div></div>
              价格可见性
            </div>
            <n-grid y-gap="25" :cols="1">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">用户组</div>
                  <div class="level-detail-div-item-contentflex">
                      <template v-if="formValue.memberGroupId">
                        <div v-for="item in formValue.memberGroupDetail" :key="item.id">{{ item.memberGroup }}</div>
                      </template>
                      <template v-else>
                        <div>全部用户</div>
                      </template>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">会员等级</div>
                  <div class="level-detail-div-item-contentflex">
                        <template v-if="formValue.memberLevelId">
                          <div v-for="item in formValue.memberLevelDetail" :key="item.id">{{ item.levelName }}</div>
                        </template>
                        <template v-else>
                          <div>全部等级</div>
                        </template>
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
  import { View } from '@/api/pmsPricePlan';
  import { State, newState } from './model';
  import { adaModalWidth } from '@/utils/hotgo';
  import {jsontoobj} from "@/utils/smjcomm";

  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref(newState(null));
  const nameLanguage = ref({
    en: '',
    zh: '',
    ja: '',
    ko: '',
    zh_CN: '',
  });
  const dialogWidth = computed(() => {
    return adaModalWidth(580);
  });

  function openModal(state: State) {
    showModal.value = true;
    loading.value = true;
    View({ id: state.id })
      .then((res) => {
        formValue.value = res;

        if(res.planShowNameLanguage){
          res.planShowNameLanguage = jsontoobj(res.planShowNameLanguage);
          nameLanguage.value.zh = res.planShowNameLanguage.zh && res.planShowNameLanguage.zh.content ? res.planShowNameLanguage.zh.content : ''
          nameLanguage.value.en = res.planShowNameLanguage.en && res.planShowNameLanguage.en.content ? res.planShowNameLanguage.en.content : ''
          nameLanguage.value.ko = res.planShowNameLanguage.ko && res.planShowNameLanguage.ko.content ? res.planShowNameLanguage.ko.content : ''
          nameLanguage.value.ja = res.planShowNameLanguage.ja && res.planShowNameLanguage.ja.content ? res.planShowNameLanguage.ja.content : ''
          if(res.planShowNameLanguage.zh_CN){
            nameLanguage.value.zh_CN = res.planShowNameLanguage.zh_CN.content ? res.planShowNameLanguage.zh_CN.content : ''
          }else if(res.planShowNameLanguage.zh_cn){
            nameLanguage.value.zh_CN = res.planShowNameLanguage.zh_cn.content ? res.planShowNameLanguage.zh_cn.content : ''
          }else{
            nameLanguage.value.zh_CN = ''
          }
        }
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
    &-content1{
      padding: 0 6px;
      height: 24px;
      line-height: 24px;
      font-weight: 400;
      font-size: 14px;
      display: inline-block;
      &.success{
        color: #19A158;
        background: #E3F4EB;
      }
      &.warning{
        color: #EFA020;
        background: #FDF1DD;
      }
    }
    &-contentflex{
      display: flex;
      align-items: center;
      div{
        font-weight: 500;
        font-size: 14px;
        color: #3D3D3D;
        line-height: 20px;
        margin-right: 15px;
      }
    }
  }
}
</style>


