export function helper(text) {
    return text.toUpperCase();
}

export function formatDate(date) {
    return date.toISOString().split('T')[0];
}
