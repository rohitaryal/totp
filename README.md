# totp

Time-based one-time password implementation from scratch

References: [How do one time passwords work?](https://dogac.dev/blog/2025/how-do-one-time-passwords-work/)

## Types of OTP algorithms

- **HOTP (HMAC-based One-time Password)** - Based on a counter that increments every time an OTP is generated [RFC 4226](https://www.rfc-editor.org/rfc/rfc4226)
- **TOTP (Time-based One-time Password)** - Based on current time (UNIX Time), typically using 30-second intervals [RFC 6238](https://www.rfc-editor.org/rfc/rfc6238)

## How to generate HOTP/TOTP?

To generate a HOTP, we need to decide on three things:

1. A secret key
2. A hash function
3. Number of digits in output

K_pad = H(K)
HMAC(K, M) = H(K_pad ^ O_pad | H(k_pad ^ I_pad | M))

Finally we can generate `HOTP` as:
HOTP(K, C) = DT(HMAC(K, C)) mod 10^digits

Replacing `C` with `c(t)`, we obtain TOTP:
TOTP(K, c(t)) = DT(HMAC(K, c(t))) mod 10^digits

From the TOTP RFC, DT is defined as:

```
    DT(String) // String = String[0]...String[19]
     Let OffsetBits be the low-order 4 bits of String[19]
     Offset = StToNum(OffsetBits) // 0 <= OffSet <= 15
     Let P = String[OffSet]...String[OffSet+3]
     Return the Last 31 bits of P
```
