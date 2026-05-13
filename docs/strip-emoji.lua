-- Emoji handling for XeLaTeX/tectonic output.
--
-- Prose text (Str): wraps unsupported chars in \emojifont so they render.
-- Inline/block code: strips them — verbatim environments can't font-switch.
--
-- Unsupported ranges in DejaVu:
--   Miscellaneous Symbols  U+2600–U+26FF  → E2 98-9B xx
--   Dingbats               U+2700–U+27BF  → E2 9C-9E xx  (e.g. ✨ U+2728)
--   Supplementary chars    U+10000+        → F0-F7 … (4-byte UTF-8)

local function utf8_char_len(byte)
  if byte < 0x80 then return 1
  elseif byte < 0xE0 then return 2
  elseif byte < 0xF0 then return 3
  else return 4 end
end

local function is_unsupported(text, i)
  local b1 = text:byte(i)
  if b1 >= 0xF0 then return 4 end          -- 4-byte: all supplementary/emoji
  if b1 == 0xE2 then
    local b2 = text:byte(i + 1) or 0
    if b2 >= 0x98 and b2 <= 0x9E then      -- Misc Symbols + Dingbats
      return 3
    end
  end
  return nil
end

-- Split text into chunks tagged plain/emoji
local function split(text)
  local chunks, plain, i = {}, "", 1
  while i <= #text do
    local emoji_len = is_unsupported(text, i)
    if emoji_len then
      if plain ~= "" then
        table.insert(chunks, {kind = "plain", text = plain})
        plain = ""
      end
      table.insert(chunks, {kind = "emoji", text = text:sub(i, i + emoji_len - 1)})
      i = i + emoji_len
    else
      local cl = utf8_char_len(text:byte(i))
      plain = plain .. text:sub(i, i + cl - 1)
      i = i + cl
    end
  end
  if plain ~= "" then table.insert(chunks, {kind = "plain", text = plain}) end
  return chunks
end

local function has_unsupported(text)
  return text:match("[\xF0-\xF7][\x80-\xBF][\x80-\xBF][\x80-\xBF]") or
         text:match("\xE2[\x98-\x9E][\x80-\xBF]")
end

-- Prose text: render emoji via \emojifont
function Str(el)
  if not has_unsupported(el.text) then return el end
  local inlines = {}
  for _, chunk in ipairs(split(el.text)) do
    if chunk.kind == "plain" then
      table.insert(inlines, pandoc.Str(chunk.text))
    else
      table.insert(inlines, pandoc.RawInline("latex",
        "{\\emojifont " .. chunk.text .. "}"))
    end
  end
  return inlines
end

-- Inline code and code blocks: strip (verbatim can't font-switch)
local function strip(text)
  text = text:gsub("[\xF0-\xF7][\x80-\xBF][\x80-\xBF][\x80-\xBF]", "")
  text = text:gsub("\xE2[\x98-\x9E][\x80-\xBF]", "")
  return text
end

function Code(el)     el.text = strip(el.text); return el end
function CodeBlock(el) el.text = strip(el.text); return el end
