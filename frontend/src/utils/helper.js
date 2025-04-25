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



export const isEmailFormatValid = (email) => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
};



export const generateUUID = () => {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
        var r = Math.random() * 16 | 0, v = c === 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
    });
};


export const invert = (hex) => {
    const rgb = hexToRgb(hex);
    const thresh = 165;
    const result = (rgb.r * 0.299 + rgb.g * 0.587 + rgb.b * 0.114) > thresh ? '#000000' : '#ffffff';
    return result;
}

const hexToRgb = (hex) => {
    const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex);
    return result ? {
        r: parseInt(result[1], 16),
        g: parseInt(result[2], 16),
        b: parseInt(result[3], 16)
    } : null;
};



export const colorToStyle = (color) => {
    const colorMap = {
        "dark": "#1A202C",
        "blue": "#3B82F6",
        "red": "#EF4444",
        "green": "#34D399",
        "yellow": "#F7DC6F",
        "indigo": "#6366F1",
        "purple": "#8B5CF6",
        "cyan": "#45A0E6",
        "gray": "#64748B",
        "lime": "#84CC16",
        "pink": "#EC4899",
        "teal": "#14B8A6"
    };
    return colorMap[color];
}
export const getColor = (percentage) => {
    if (percentage >= 90) {
        return "green";
    } else if (percentage >= 75) {
        return "lime";
    } else if (percentage >= 50) {
        return "yellow";
    } else if (percentage >= 25) {
        return "cyan";
    } else if (percentage > 0) {
        return "red";
    } else {
        return "dark";
    }
};



export const daysToMilliseconds = (days) => {
    return days * 24 * 60 * 60 * 1000;
}

export const toSnakeCase = (str) => {
    return str
        .toLowerCase()
        .replace(/[^\w\s]/g, '')
        .split(/\s+/)
        .join('_');
}

export const money = (val, friction = 2) => {
    if (!val) return 0;
    return val.toLocaleString('id-ID', { useGrouping: true, maximumFractionDigits: friction });
}

export const nl2br = (str) => {
    return str.replace(/(\r\n|\n\r|\r|\n)/g, '<br>');
}


export const randomString = (length) => {
    let result = '';
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    const charactersLength = characters.length;
    for (let i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
}



export const getContrastColor = (hex) => {
    if (!hex) return "#000000";
    const color = hex.replace(/^#/, '');
    const r = parseInt(color.substr(0, 2), 16);
    const g = parseInt(color.substr(2, 2), 16);
    const b = parseInt(color.substr(4, 2), 16);
    const yiq = (r * 299 + g * 587 + b * 114) / 1000;
    return (yiq >= 128) ? '#000000' : '#FFFFFF';
}


export const randomColor = ({ luminosity = "dark" } = {}) => {
    const randomHex = () => {
        const hexPart = Math.floor(Math.random() * 256).toString(16);
        return hexPart.length === 1 ? '0' + hexPart : hexPart;
    };

    let color = `#${randomHex()}${randomHex()}${randomHex()}`;

    if (luminosity === "dark") {
        while (getContrastColor(color) !== '#FFFFFF') {
            color = `#${randomHex()}${randomHex()}${randomHex()}`;
        }
    }

    return color;
};

