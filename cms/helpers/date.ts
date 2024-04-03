import moment from "moment"

export const formatDate = (value: string, type : string = "DD-MM-YYYY") => {
   return moment(value).format(type)
}