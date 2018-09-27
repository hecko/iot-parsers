
import binascii
import struct
from pprint import pprint


def parse_sigfox_packet(data):
    out = {}
    out["temp"] = struct.unpack(">h", binascii.unhexlify(data[0:4]))[0] / 16.0
    out["humi"] = struct.unpack(">B", binascii.unhexlify(data[6:8]))[0]
    out["batt"] = struct.unpack(">H", binascii.unhexlify(data[-4:]))[0]
    return out


if __name__ == "__main__":
    PAYLOADS = ("018100280266", "FEFC00440258")
    for payload in PAYLOADS:
        print("Parsing Sigfox packet: {}".format(payload))
        pprint(parse_sigfox_packet(payload))
