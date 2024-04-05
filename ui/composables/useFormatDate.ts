export default function (value : string) {
    let date = new Date(value)
    return date.toLocaleDateString('en-GB')
}