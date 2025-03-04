export function debounce(func, wait, immediate) {
    let timeout;
    return function (...args) {
        const context = this;
        const later = function () {
            timeout = null;
            if (!immediate) func.apply(context, args);
        };
        const callNow = immediate && !timeout;
        clearTimeout(timeout);
        timeout = setTimeout(later, wait);
        if (callNow) func.apply(context, args);
    };
}

export const getPagination = (resp) => {
    const { page, size, max_page, total_pages, total, last, first, visible } = resp;
    return { page, size, max_page, total_pages, total, last, first, visible }
}

export const initial = (str) => {
    if (!str) return '';
    return str.split(' ').map(word => word.charAt(0).toUpperCase()).join('');
};

