import { format } from 'date-fns';

// 将时间格式化为 'yyyy-MM-dd HH:mm:ss'
export function formatTime(time: number) {
  const date = new Date(time);
  return format(date, 'yyyy-MM-dd HH:mm:ss');
}
