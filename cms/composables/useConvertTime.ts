import moment from "moment";
export default async function (value: string, type: string) {
  return await moment(value).utc().format(type)
}
