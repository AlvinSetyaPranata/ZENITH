export function convertTimestampToDate(timestamp: string) {
    const date = new Date(timestamp)

    const hours = date.getUTCHours()

    return `${(hours % 12 || 12).toString().padStart(2, '0')}:${date.getUTCMinutes().toString().padStart(2, '0')} ${hours >= 12 ? "AM" : "PM"}`
}