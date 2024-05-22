import request from "@/utils/request";

const BASE_API = "/admin/feat";

export default {
  all(query) {
    return request({
      url: BASE_API + '/{{ . }}/all',
      method: "get",
      params: query,
    });
  },
  page(query) {
    return request({
      url: BASE_API + '/{{ . }}/page',
      method: "get",
      params: query,
    });
  },
  create(data) {
    return request({
      url: BASE_API + '/{{ . }}/create',
      method: "post",
      data,
    });
  },
  update(data) {
    return request({
      url: BASE_API + '/{{ . }}/update',
      method: "post",
      data,
    });
  },
  detail(id) {
      return request({
        url: BASE_API + '/{{ . }}/detail',
        method: "get",
        params: { id: id },
      });
  },
  delete(id) {
    return request({
      url: BASE_API + '/{{ . }}/delete',
      method: "post",
      data: { id: id },
    });
  },
  deletes(ids) {
    return request({
      url: BASE_API + '/{{ . }}/deletes',
      method: "post",
      data: { ids: ids },
    });
  },
};