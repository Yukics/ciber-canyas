export function parseEmojiToHtml(emoji:string): string{
    return String.fromCodePoint(emoji.replace('U+','0x'))
}